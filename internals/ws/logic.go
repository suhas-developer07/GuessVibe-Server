package ws

import (
	"encoding/json"
	"log"

	wsmodels "github.com/suhas-developer07/GuessVibe-Server/internals/models/Websocket_model"
)

func HandleIncomingMessage(c *Client, msg []byte){
	var input wsmodels.ClientMessage

	if err := json.Unmarshal(msg,&input);err!=nil{
		log.Println("invalid ws msg",err)
		return 
	}

	switch input.Type{
	case "init":
	//TODO: here i should handle init message from client when he connects first time to socket 
	// send this request to redis and grpc server
	log.Println("First question from the server when he gets connected to a server")
	HandleInit(c,input)

	case "answer":
		//user answered a questino
		//TODO:send this request to grpc server and get the next question
		HandleAnswer(c,input)
	}
}

func SendQuestion(c *Client, question string, num int) {
	resp := wsmodels.ServerMessage{
		Type:           "question",
		Question:       question,
		QuestionNumber: num,
	}

	b, _ := json.Marshal(resp)
	c.Send <- b
}

func HandleInit(c *Client, msg wsmodels.ClientMessage){
	//First question from the LLM

	firstQuestion := "Is it alive"

	SendQuestion(c,firstQuestion,1)
}

func HandleAnswer(c *Client, msg wsmodels.ClientMessage){
	nextQ := "It is bigger than football"
	SendQuestion(c,nextQ,2)
}