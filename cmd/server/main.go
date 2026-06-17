package main

import (
	"log"

	"github.com/akash/go-user-api/config"
	"github.com/akash/go-user-api/internal/handler"
	"github.com/akash/go-user-api/internal/logger"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {

	err := logger.InitLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Log.Sync()

	err = config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server Running")
	})

	app.Post("/users", handler.CreateUser)
	app.Get("/users", handler.GetAllUsers)
	app.Get("/users/:id", handler.GetUserByID)
	app.Put("/users/:id", handler.UpdateUser)
	app.Delete("/users/:id", handler.DeleteUser)

	logger.Log.Info(
		"Server Started",
		zap.String("port", "3000"),
	)

	log.Fatal(app.Listen(":3000"))
}
