package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
	grpcclient "github.com/suhas-developer07/GuessVibe-Server/internals/grpc_client"
	"github.com/suhas-developer07/GuessVibe-Server/internals/router"
	"github.com/suhas-developer07/GuessVibe-Server/internals/session"
	"github.com/suhas-developer07/GuessVibe-Server/internals/ws"
)

func main() {
	e := echo.New()

	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Warning unable to find .env file")
	}

	Redis_url := os.Getenv("REDIS_URL")
	if Redis_url == "" {
		log.Fatal("Redis_url not found in environment variables")
	}
	
	opt,_ := redis.ParseURL(Redis_url)
	rbd := redis.NewClient(opt)

	repo := session.NewRedisRepo(rbd)
	svc := session.NewService(repo)
	ws.InjectSessionService(svc)

	llmHost := os.Getenv("LLM_HOST") // e.g. "python-service.up.railway.internal"
	llmPort := os.Getenv("LLM_PORT")

	if llmHost == "" || llmPort == "" {
		log.Fatal("LLM_HOST or LLM_PORT not found in environment variables")
	}

	llmClient := grpcclient.NewLLMClient(fmt.Sprintf("%s:%s",llmHost,llmPort))

	client := Connect()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			echo.GET,
			echo.POST,
			echo.PUT,
			echo.DELETE,
			echo.PATCH,
			echo.OPTIONS,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
	}))

	hub := ws.NewHub()
	go hub.Run()

	router.LoadHTTPRoutes(e, client)
	router.LoadWSRoutes(e, hub, llmClient)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := e.Start(":"+port); err != nil {
		e.Logger.Fatal("Shutting down the server")
	}
}
