package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/vpovarna/go-todo-api/config"
	"github.com/vpovarna/go-todo-api/db"
	"go.uber.org/fx"
)

func main() {
	todoServiceConfig := config.LoadEnv()
	_ = db.CreateMySQLConnection(todoServiceConfig)

	//fx.New(
	//	fx.Provide(),
	//	fx.Invoke(newFiberServer),
	//).Run()
}

func newFiberServer(lc fx.Lifecycle) *fiber.App {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

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
