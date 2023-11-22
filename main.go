package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/vpovarna/go-todo-api/config"
	"github.com/vpovarna/go-todo-api/db"
	"github.com/vpovarna/go-todo-api/handlers"
	"github.com/vpovarna/go-todo-api/repository"
	"github.com/vpovarna/go-todo-api/router"
	"go.uber.org/fx"
)

func main() {
	fx.New(fx.Provide(
		config.LoadEnv,
		db.CreateMySQLConnection,
		repository.NewTodoStorage,
		handlers.NewHealthCheckHandler,
		handlers.NewTodoHandlers,
	),
		fx.Invoke(newFiberServer),
	).Run()
}

func newFiberServer(lc fx.Lifecycle, todoHandler *handlers.TodoHandlers, healthCheckHandler *handlers.HealthCheckHandler) *fiber.App {

	app := fiber.New()

	//TODO: double check the middleware
	app.Use(logger.New())

	router.SetupRoutes(app, healthCheckHandler, todoHandler)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Starting fiber on port: 18081")
			//TODO: read the port from an env variable via config
			go func() {
				err := app.Listen(":18081")
				if err != nil {
					log.Panic("Unable to start server. Error: ", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})
	return app
}
