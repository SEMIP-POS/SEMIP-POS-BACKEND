// internal/repository/postgresql/postgresql.go
package postgresqlRepository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IPostgres interface {
	Close() error
	Ping() error
	GetDB() *gorm.DB
	WithContext(ctx context.Context) *gorm.DB
	Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error
}

type Config struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DBName       string
	MaxIdleConns int
	MaxIdleTime  int
	MaxLifeTime  int
	MaxOpenConns int
}

type Postgres struct {
	DB *gorm.DB
}

func New(config Config) (IPostgres, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Host,
		config.Username,
		config.Password,
		config.DBName,
		config.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	setDBConfig(sqlDB, config)

	return &Postgres{DB: db}, nil
}

func setDBConfig(sqlDB *sql.DB, config Config) {
	if config.MaxIdleConns == 0 {
		config.MaxIdleConns = 15
	}

	if config.MaxOpenConns == 0 {
		config.MaxOpenConns = 25
	}

	if config.MaxIdleTime == 0 {
		config.MaxIdleTime = 300 // 5 Mins
	}

	if config.MaxLifeTime == 0 {
		config.MaxLifeTime = 300 // 5 Mins
	}

	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetConnMaxIdleTime(time.Duration(config.MaxIdleTime) * time.Second)
	sqlDB.SetConnMaxLifetime(time.Duration(config.MaxLifeTime) * time.Second)
}

func (p *Postgres) Close() error {
	sqlDB, err := p.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (p *Postgres) Ping() error {
	sqlDB, err := p.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

func (p *Postgres) Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return p.DB.WithContext(ctx).Transaction(fn)
}

func (p *Postgres) GetDB() *gorm.DB {
	return p.DB
}

func (p *Postgres) WithContext(ctx context.Context) *gorm.DB {
	return p.DB.WithContext(ctx)
}
