package healthCheck

import (
	postgresqlRepository "github.com/SEMIP-POS/semip-pos-backend/internal/repository/postgresql"

	"github.com/SEMIP-POS/semip-pos-backend/internal/repository"
)

type healthCheck struct {
	db postgresqlRepository.IPostgres
}

func NewHealthRepository(
	db postgresqlRepository.IPostgres,
) repository.IHealthCheckRepository {
	return &healthCheck{
		db: db,
	}
}
