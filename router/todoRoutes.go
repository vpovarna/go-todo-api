package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vpovarna/go-todo-api/handlers"
)

func SetupRoutes(app *fiber.App, healthCheckHandler *handlers.HealthCheckHandler, todoHandlers *handlers.TodoHandlers) *fiber.App {

	app.Get("/health", healthCheckHandler.HealthCheck)

	todoApis := app.Group("/api/todos")

	todoApis.Get("/:id", todoHandlers.GetTodoById)
	todoApis.Post("/", todoHandlers.CreateTodo)
	todoApis.Post("/:id", todoHandlers.CompleteTodo)
	todoApis.Delete("/:id", todoHandlers.DeleteTodo)
	todoApis.Get("/", todoHandlers.GetAllTodos)

	return app
}
