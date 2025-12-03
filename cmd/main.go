package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
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

	err := godotenv.Load(".env")
	MongoDb := os.Getenv("MONGODB_URI")
	log.Println("MONGODB_URI: ", MongoDb)
	if MongoDb == "" {
		log.Fatal("MONGODB_URI not found in environment variables")
	}

	if err != nil {
		log.Println("Warning unable to find .env file")
	}

	llmClient := grpcclient.NewLLMClient("localhost:50051")

	client := Connect()

	hub := ws.NewHub()
	go hub.Run()

	router.LoadHTTPRoutes(e, client)
	router.LoadWSRoutes(e, hub, llmClient)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal("Shutting down the server")
	}
}
