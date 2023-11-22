package handlers

import "github.com/gofiber/fiber/v2"

type HealthCheckHandler struct {
}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (*HealthCheckHandler) HealthCheck(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{"status": "ok"})
}
