package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/vpovarna/go-todo-api/config"
	"github.com/vpovarna/go-todo-api/db"
	"github.com/vpovarna/go-todo-api/handlers"
	"github.com/vpovarna/go-todo-api/repository"
	"github.com/vpovarna/go-todo-api/router"
)

func main() {
	ctx := context.Background()

	todoServiceConfig := config.LoadEnv()
	conn := db.CreateMySQLConnection(ctx, todoServiceConfig)
	todoStorage := repository.NewTodoStorage(conn)
	handler := handlers.NewTodoHandlers(todoStorage)

	app := fiber.New()
	router.SetupRoutes(app, handler)
	// TODO: Move the listen address to env file
	log.Fatal(app.Listen(":18081"))
}
