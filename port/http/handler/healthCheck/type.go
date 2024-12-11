package healthCheckHandler

import (
	healthCheckService "github.com/SEMIP-POS/semip-pos-backend/internal/service"
)

type HealthCheckHandler struct {
	healthCheckSvc healthCheckService.IHealthService
}

func NewHealthHandler(
	healthCheckSvc healthCheckService.IHealthService,
) *HealthCheckHandler {
	return &HealthCheckHandler{
		healthCheckSvc: healthCheckSvc,
	}

}
