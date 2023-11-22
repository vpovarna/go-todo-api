package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vpovarna/go-todo-api/handlers"
)

func SetupRoutes(app *fiber.App, handlers *handlers.TodoHandlers) {
	api := app.Group("/api")

	api.Get("/:id", handlers.GetTodoById)
	api.Post("/", handlers.CreateTodo)
	api.Put("/:id", handlers.CompleteTodo)
	api.Delete("/:id", handlers.DeleteTodo)
}
