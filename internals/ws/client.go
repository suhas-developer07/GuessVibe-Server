package ws

import (
	"log"

	"github.com/gorilla/websocket"
	grpcclient "github.com/suhas-developer07/GuessVibe-Server/internals/grpc_client"
)

type Client struct {
	Hub       *Hub
	Conn      *websocket.Conn
	Send      chan []byte
	LLM       *grpcclient.LLMClient
	SessionID string
	UserID    string
}

// ReadPump handles incoming messages from frontend
func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}

		// send message to your WS handler logic:
		HandleIncomingMessage(c, msg)
	}
}

// WritePump handles outgoing messages to frontend
func (c *Client) WritePump() {
	defer c.Conn.Close()

	for msg := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("write error:", err)
			return
		}
	}
}
