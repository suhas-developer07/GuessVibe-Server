package main

import (
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	grpcclient "github.com/suhas-developer07/GuessVibe-Server/internals/grpc_client"
	"github.com/suhas-developer07/GuessVibe-Server/internals/router"
	"github.com/suhas-developer07/GuessVibe-Server/internals/session"
	"github.com/suhas-developer07/GuessVibe-Server/internals/ws"
)

func main() {
	e := echo.New()

	rbd := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	repo := session.NewRedisRepo(rbd)
	svc := session.NewService(repo)
	ws.InjectSessionService(svc)

	llmClient := grpcclient.NewLLMClient("localhost:50051")


	hub:= ws.NewHub()
	go hub.Run()

	router.LoadHTTPRoutes(e)
	router.LoadWSRoutes(e, hub,llmClient)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal("Shutting down the server")
	}
}