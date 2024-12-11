package service

import (
	"context"

	healthCheckDomain "github.com/SEMIP-POS/semip-pos-backend/internal/domain/healthCheck"
)

type IHealthService interface {
	Check(ctx context.Context) (*healthCheckDomain.Health, error)
}
