package handler

import (
	"github.com/gofiber/fiber/v2"
)

type V1HealthCheckHandler interface {
	Check(c *fiber.Ctx) error
}
