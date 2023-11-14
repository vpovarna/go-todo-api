package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(),
		fx.Invoke(newFiberServer),
	).Run()
}

func newFiberServer(lc fx.Lifecycle) *fiber.App {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	_ = app.Group("/todo")
	//todoGroup.Get("/:id", handlers.CreateTodo)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			//TODO: Use env variables
			fmt.Println("Starting http server on port 18081")
			go app.Listen(":18081")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	return app
}
