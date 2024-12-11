package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/SEMIP-POS/semip-pos-backend/config"
	postgresqlRepository "github.com/SEMIP-POS/semip-pos-backend/internal/repository/postgresql"
	healthCheckRepo "github.com/SEMIP-POS/semip-pos-backend/internal/repository/postgresql/healthCheck"
	healthCheckService "github.com/SEMIP-POS/semip-pos-backend/internal/service/healthCheck"
	route "github.com/SEMIP-POS/semip-pos-backend/port/http"
	healthCheckHandler "github.com/SEMIP-POS/semip-pos-backend/port/http/handler/healthCheck"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveHttp)
}

var serveHttp = &cobra.Command{
	Use:   "serveHttp",
	Short: "start HTTP server",
	Long:  "Start HTTP server with configured settings",
	Run: func(cmd *cobra.Command, args []string) {
		// Load config files
		cfg, secret, err := config.LoadConfig(cfgFile, scrtFile)
		if err != nil {
			log.Fatalf("Failed to load configuration: %v", err)
		}

		// Initialize database
		db, err := postgresqlRepository.New(postgresqlRepository.Config{
			Host:         cfg.DatabaseConfig.Host,
			Port:         cfg.DatabaseConfig.Port,
			Username:     secret.DatabaseSecret.Username,
			Password:     secret.DatabaseSecret.Password,
			DBName:       secret.DatabaseSecret.Database,
			MaxIdleConns: cfg.DatabaseConfig.MaxIdleConns,
			MaxOpenConns: cfg.DatabaseConfig.MaxOpenConns,
			MaxIdleTime:  cfg.DatabaseConfig.MaxIdleTime,
			MaxLifeTime:  cfg.DatabaseConfig.MaxLifeTime,
		})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		// Initialize Fiber app
		app := fiber.New(fiber.Config{
			ReadTimeout:  cfg.ServerConfig.ReadTimeout,
			WriteTimeout: cfg.ServerConfig.WriteTimeout,
			AppName:      cfg.ServiceName,
		})

		// Initialize repositories
		healthRepo := healthCheckRepo.NewHealthRepository(db)

		// Initialize services
		healthService := healthCheckService.NewHealthService(healthRepo, cfg)

		// Initialize handlers
		healthHandler := healthCheckHandler.NewHealthHandler(healthService)

		// Setup routes
		routeConfig := &route.RouteConfig{
			App:           app,
			HealthHandler: healthHandler,
		}
		route.Setup(routeConfig)

		// Start server
		go func() {
			if err := app.Listen(fmt.Sprintf(":%s", cfg.Port)); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		log.Printf("Server started on port %s", cfg.Port)

		// Wait for interrupt signal
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		// Shutdown server
		if err := app.Shutdown(); err != nil {
			log.Fatalf("Server Shutdown Failed:%+v", err)
		}

		log.Println("Server stopped")
	},
}
