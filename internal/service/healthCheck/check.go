package healthCheckService

import (
	"context"

	healthCheckDomain "github.com/SEMIP-POS/semip-pos-backend/internal/domain/healthCheck"
)

func (s *healthService) Check(ctx context.Context) (*healthCheckDomain.Health, error) {
	err := s.repo.CheckDB(ctx)
	dbStatus := "up"
	if err != nil {
		dbStatus = "down"
	}

	return &healthCheckDomain.Health{
		Status:      "ok",
		ServiceName: s.config.ServiceName,
		Version:     s.config.ServiceVersion,
		Database:    dbStatus,
	}, nil
}
