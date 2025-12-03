package ws

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/suhas-developer07/GuessVibe-Server/generated/proto"
	//models "github.com/suhas-developer07/GuessVibe-Server/internals/models/RedisSession_model"
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

func SendFinalGuess(c *Client, guess *pb.LLMGuessResponse) {
	resp := wsmodels.ServerMessage{
		Type:         "final_guess",
		Guess:        guess.Guess,
		Confidence:   guess.Confidence,
		Alternatives: guess.Alternatives,
		SessionID:    c.SessionID,
	}

	b, _ := json.Marshal(resp)
	c.Send <- b
}

func HandleInit(c *Client, msg wsmodels.ClientMessage) {
	ctx := context.Background()

	// 1. Create new session in Redis
	state, err := sessionService.CreateSession(ctx, msg.UserID)
	if err != nil {
		log.Println("session create failed:", err)
		return
	}

	// 2. Convert to proto for gRPC call
	protoState := state.ToProto()

	// 3. Call Python LLM for FIRST QUESTION
	firstQuestion, err := c.LLM.GenerateFirstQuestion(protoState)
	if err != nil {
		log.Println("LLM error:", err)
		return
	}

	// 4. Update local game state
	state.QuestionCount = 1
	state.History = append(state.History, session.QA{
		Question: firstQuestion,
		Answer:   "",
	})

	// 5. Save back to Redis
	if err := sessionService.Repo.UpdateSession(ctx, state); err != nil {
		log.Println("redis update failed:", err)
	}

	// 6. Attach session to WS client
	c.SessionID = state.SessionID
	c.UserID = state.UserID

	// 7. Send question to frontend
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

	if len(state.History) == 0 {
		log.Println("ERROR: No previous question to attach answer to")
		return
	}

	// 2. Update last question with user answer
	state.History[len(state.History)-1].Answer = msg.Answer

	// 3. DECIDE: next question OR final guess?
	if state.QuestionCount >= 15 {
		// ------------ FINAL GUESS ------------ //

		protoState := state.ToProto()
		guessResp, err := c.LLM.GenerateFinalGuess(protoState)
		if err != nil {
			log.Println("final guess error:", err)
			return
		}

		// Send guess to frontend
		SendFinalGuess(c, guessResp)
		SendSessionEnd(c)

		// Close session in Redis
		_ = sessionService.Repo.DeleteSession(ctx, state.SessionID)

		CloseClient(c)

		return
	}

	// ------------ NEXT QUESTION ------------ //

	// Increase count BEFORE calling LLM
	state.QuestionCount++

	protoState := state.ToProto()

	nextQuestion, err := c.LLM.GenerateNextQuestion(protoState)
	if err != nil {
		log.Println("LLM next question error:", err)
		return
	}

	// Add new Q to history
	state.History = append(state.History, session.QA{
		Question: nextQuestion,
		Answer:   "",
	})

	// Save updated session
	if err := sessionService.Repo.UpdateSession(ctx, state); err != nil {
		log.Println("redis update failed:", err)
		return
	}

	// Send next question
	SendQuestion(c, state.SessionID, nextQuestion, state.QuestionCount)
}

func SendSessionEnd(c *Client) {
	resp := wsmodels.ServerMessage{
		Type: "session_end",
	}

	b, _ := json.Marshal(resp)
	c.Send <- b
}
