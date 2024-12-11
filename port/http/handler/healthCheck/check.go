package healthCheckHandler

import (
	"github.com/gofiber/fiber/v2"
)

func (h *HealthCheckHandler) Check(c *fiber.Ctx) error {
	health, err := h.healthCheckSvc.Check(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	return c.JSON(health)
}
