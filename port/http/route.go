package route

import (
	handler "github.com/SEMIP-POS/semip-pos-backend/port/http/handler"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App           *fiber.App
	HealthHandler handler.V1HealthCheckHandler
}

func Setup(config *RouteConfig) {
	config.App.Get("/health", config.HealthHandler.Check)
}
