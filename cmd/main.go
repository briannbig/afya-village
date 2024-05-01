package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/briannbig/afya-village/internal/application/controller"
	"github.com/briannbig/afya-village/internal/application/middleware"
	"github.com/briannbig/afya-village/internal/application/service"
	"github.com/briannbig/afya-village/internal/domain/queue"
	"github.com/briannbig/afya-village/internal/infra/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/nats-io/nats.go"
)

var (
	userController controller.UserController
	emailNotifier  service.Notifier
)

func main() {

	log.Printf("<<<<<<<<<<<<<<<Afya Village>>>>>>>>>>>>>>>")

	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading env file", envErr)
	}

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Printf("could not connect to queue --- %s", err.Error())
		os.Exit(1)
	}

	q := queue.New(nc)

	db := database.New()

	userController = controller.NewUserController(&db, &q)

	emailNotifier = service.Emailnotifier()
	q.RegisterConsumer(queue.EventUserCreated, emailNotifier.Notify)

	r := router()

	r.Run(":5050")

}

func router() *gin.Engine {

	server := gin.Default()

	setupLogger()

	server.Use(gin.Recovery(), middleware.Logger())

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to Afya Village"})
	})

	userRoutes := server.Group("/users")
	{
		userRoutes.POST("/register", userController.Register)
	}
	server.GET("/users/register", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to Afya Village"})
	})

	return server
}

func setupLogger() {
	f, _ := os.Create("app-server.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
