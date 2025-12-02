package ws

import (
	"context"
	"encoding/json"
	"log"

	models "github.com/suhas-developer07/GuessVibe-Server/internals/models/RedisSession_model"
	wsmodels "github.com/suhas-developer07/GuessVibe-Server/internals/models/Websocket_model"
	"github.com/suhas-developer07/GuessVibe-Server/internals/session"
)

var sessionService *session.Service // injected from main.go

func InjectSessionService(s *session.Service) {
	sessionService = s
}

func HandleIncomingMessage(c *Client, msg []byte) {
	var input wsmodels.ClientMessage

	if err := json.Unmarshal(msg, &input); err != nil {
		log.Println("invalid ws msg", err)
		return
	}

	switch input.Type {
	case "init":
		//TODO: here i should handle init message from client when he connects first time to socket
		// send this request to redis and grpc server
		log.Println("First question from the server when he gets connected to a server")
		HandleInit(c, input)

	case "answer":
		//user answered a questino
		//TODO:send this request to grpc server and get the next question
		HandleAnswer(c, input)
	}
}

func SendQuestion(c *Client, sessionID string, question string, num int) {
	resp := wsmodels.ServerMessage{
		Type:           "question",
		SessionID:      sessionID,
		Question:       question,
		QuestionNumber: num,
	}

	b, _ := json.Marshal(resp)
	c.Send <- b
}

func HandleInit(c *Client, msg wsmodels.ClientMessage) {
    ctx := context.Background()

    // 1. Create new session
    state, err := sessionService.CreateSession(ctx, msg.UserID)
    if err != nil {
        log.Println("session create failed:", err)
        return
    }

    // TEMP FIRST QUESTION
    firstQuestion := "Is it alive?"

    // ❗️ IMPORTANT: Add first question to history
    state.History = append(state.History, models.QA{
        Question: firstQuestion,
        Answer:   "",
    })

    state.QuestionCount = 1

    // Save back to Redis
    if err := sessionService.Repo.UpdateSession(ctx, state); err != nil {
        log.Println("failed to update session:", err)
    }

    // Attach to client
    c.SessionID = state.SessionID
    c.UserID = state.UserID

    // Send first question
    SendQuestion(c, state.SessionID, firstQuestion, state.QuestionCount)
}
func HandleAnswer(c *Client, msg wsmodels.ClientMessage) {
    ctx := context.Background()

    // 1. Load session
    state, err := sessionService.Repo.GetSession(ctx, msg.SessionID)
    if err != nil {
        log.Println("session not found:", err)
        return
    }

    // SAFETY CHECK — history must not be empty
    if len(state.History) == 0 {
        log.Println("ERROR: history empty, cannot append answer")
        return
    }

    // 2. Update last Q with answer
    state.History[len(state.History)-1].Answer = msg.Answer

    // 3. Increase question count
    state.QuestionCount++

    // ❗️Create a mock next question
    nextQuestion := "Is it bigger than a football?"

    // 4. Append new question placeholder to history
    state.History = append(state.History, models.QA{
        Question: nextQuestion,
        Answer:   "",
    })

    // Save back to Redis
    if err := sessionService.Repo.UpdateSession(ctx, state); err != nil {
        log.Println("failed to update session:", err)
        return
    }

    // 5. Send next question
    SendQuestion(c, state.SessionID, nextQuestion, state.QuestionCount)
}
