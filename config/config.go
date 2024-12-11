// internal/config/config.go
package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Environment    string `mapstructure:"ENVIRONMENT"`
	ServiceName    string `mapstructure:"SERVICE_NAME"`
	ServiceVersion string `mapstructure:"SERVICE_VERSION"`
	Port           string `mapstructure:"PORT"`

	DatabaseConfig DatabaseConfig `mapstructure:"DATABASE"`
	JWTConfig      JWTConfig      `mapstructure:"JWT"`
	ServerConfig   ServerConfig   `mapstructure:"SERVER"`
}

type Secret struct {
	DatabaseSecret DatabaseSecret `mapstructure:"DATABASE"`
	JWTSecret      JWTSecret      `mapstructure:"JWT"`
}

type DatabaseConfig struct {
	Dialect      string `mapstructure:"DIALECT"`
	Host         string `mapstructure:"HOST"`
	Port         string `mapstructure:"PORT"`
	MaxIdleConns int    `mapstructure:"MAX_IDLE_CONNS"`
	MaxOpenConns int    `mapstructure:"MAX_OPEN_CONNS"`
	MaxIdleTime  int    `mapstructure:"MAX_IDLE_TIME"`
	MaxLifeTime  int    `mapstructure:"MAX_LIFE_TIME"`
}

type DatabaseSecret struct {
	Database string `mapstructure:"DB_NAME"`
	Username string `mapstructure:"USERNAME"`
	Password string `mapstructure:"PASSWORD"`
}

type JWTConfig struct {
	ExpirationTime time.Duration `mapstructure:"EXPIRATION_TIME"`
}

type JWTSecret struct {
	SecretKey string `mapstructure:"SECRET_KEY"`
}

type ServerConfig struct {
	ReadTimeout  time.Duration `mapstructure:"READ_TIMEOUT"`
	WriteTimeout time.Duration `mapstructure:"WRITE_TIMEOUT"`
}

var env string

func LoadConfig(configPath, secretPath string) (*Config, *Secret, error) {
	// Load Config
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, nil, fmt.Errorf("error reading config file: %w", err)
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	viper.Reset()

	// Load Secret
	viper.SetConfigFile(secretPath)
	var secret Secret
	if err := viper.ReadInConfig(); err != nil {
		return nil, nil, fmt.Errorf("error reading secret file: %w", err)
	}
	if err := viper.Unmarshal(&secret); err != nil {
		return nil, nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	env = config.Environment

	return &config, &secret, nil
}

// GetDBConnString returns the formatted database connection string
func (c *Config) GetDBConnString(s *DatabaseSecret) string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		c.DatabaseConfig.Dialect,
		s.Username,
		s.Password,
		c.DatabaseConfig.Host,
		c.DatabaseConfig.Port,
		s.Database,
	)
}
