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
    closed    bool // prevent double unregister
}

// ReadPump handles incoming messages from frontend
func (c *Client) ReadPump() {
    defer CloseClient(c)

    for {
        _, msg, err := c.Conn.ReadMessage()
        if err != nil {
            log.Println("read error:", err)
            break
        }

        HandleIncomingMessage(c, msg)
    }
}


// WritePump handles outgoing messages to frontend
func (c *Client) WritePump() {
    defer c.Conn.Close()

    for msg := range c.Send {
        if err := c.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
            log.Println("write error:", err)
            return
        }
    }
}

func CloseClient(c *Client) {
    if c.closed {
        return
    }
    c.closed = true

    // Ask Hub to unregister (Hub closes channel)
    c.Hub.Unregister <- c

    // Close WebSocket connection
    _ = c.Conn.Close()

    log.Println("Client disconnected cleanly after final guess")
}
