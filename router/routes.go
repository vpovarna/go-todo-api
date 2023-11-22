package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vpovarna/go-todo-api/handlers"
)

func SetupRoutes(app *fiber.App, handlers *handlers.TodoHandlers) {
	todoApis := app.Group("/api/todos")

	todoApis.Get("/:id", handlers.GetTodoById)
	todoApis.Post("/", handlers.CreateTodo)
	todoApis.Post("/:id", handlers.CompleteTodo)
	todoApis.Delete("/:id", handlers.DeleteTodo)
	todoApis.Get("/", handlers.GetAllTodos)
}
