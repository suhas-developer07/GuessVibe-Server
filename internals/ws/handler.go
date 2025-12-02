package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	grpcclient "github.com/suhas-developer07/GuessVibe-Server/internals/grpc_client"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebsocketHandler(hub *Hub,llm *grpcclient.LLMClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}

		log.Println("Client get coneected")

		client := &Client{
			Hub:  hub,
			Conn: conn,
			Send: make(chan []byte, 256),
			LLM:llm,
		}

		client.Hub.Register <- client

		go client.WritePump()
		go client.ReadPump()

		return nil
	}
}