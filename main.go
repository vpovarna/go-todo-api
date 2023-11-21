package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vpovarna/go-todo-api/config"
	"github.com/vpovarna/go-todo-api/db"
	"github.com/vpovarna/go-todo-api/handlers"
	"github.com/vpovarna/go-todo-api/router"
	"github.com/vpovarna/go-todo-api/storage"
)

func main() {
	todoServiceConfig := config.LoadEnv()
	conn := db.CreateMySQLConnection(todoServiceConfig)
	todoStorage := storage.NewTodoStorage(conn)
	handler := handlers.NewTodoHandlers(todoStorage)

	app := fiber.New()
	router.SetupRoutes(app, handler)
	app.Listen(":18081")
}
