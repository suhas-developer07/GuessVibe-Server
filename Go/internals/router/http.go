package router

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/suhas-developer07/GuessVibe-Server/internals/handlers"
	"github.com/suhas-developer07/GuessVibe-Server/internals/repository"
	userService "github.com/suhas-developer07/GuessVibe-Server/internals/service/user"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoadHTTPRoutes(e *echo.Echo, client *mongo.Client) {
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	// Initialize dependencies
	ctx := context.Background()
	err := godotenv.Load(".env")
	if err != nil {
		// handle error
	}
	databaseName := os.Getenv("DATABASE_NAME")
	db := client.Database(databaseName)
	repo := &repository.MongoRepo{
		Db:  db,
		Ctx: ctx,
	}
	userSvc := userService.NewUserService(repo)
	userHandler := &handlers.UserHandler{
		UserServices: userSvc,
	}

	// User routes
	e.POST("/register", userHandler.UserRegisterHandler)
	e.POST("/login", userHandler.UserLoginHandler)
	e.POST("/logout", userHandler.LogoutUserHandler)
}
