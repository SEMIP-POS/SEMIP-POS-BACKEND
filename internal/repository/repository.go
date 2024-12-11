package repository

import (
	"context"
)

type IHealthCheckRepository interface {
	CheckDB(ctx context.Context) error
}
