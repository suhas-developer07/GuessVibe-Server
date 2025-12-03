package router

import (
	"github.com/labstack/echo/v4"
	grpcclient "github.com/suhas-developer07/GuessVibe-Server/internals/grpc_client"
	"github.com/suhas-developer07/GuessVibe-Server/internals/ws"
)

func LoadWSRoutes(e *echo.Echo,hub *ws.Hub,llm *grpcclient.LLMClient) {
	e.GET("/ws",ws.WebsocketHandler(hub,llm))
}