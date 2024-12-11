package healthCheckService

import (
	"github.com/SEMIP-POS/semip-pos-backend/config"
	"github.com/SEMIP-POS/semip-pos-backend/internal/repository"
)

type healthService struct {
	repo   repository.IHealthCheckRepository
	config *config.Config
}

func NewHealthService(repo repository.IHealthCheckRepository, config *config.Config) *healthService {
	return &healthService{
		repo:   repo,
		config: config,
	}
}
