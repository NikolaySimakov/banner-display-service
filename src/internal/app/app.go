package app

import (
	"banner-display-service/src/config"
	v1 "banner-display-service/src/internal/controller/http/v1"
	"banner-display-service/src/internal/repositories"
	"banner-display-service/src/internal/services"
	"banner-display-service/src/pkg/httpserver"
	"banner-display-service/src/pkg/postgres"
	"banner-display-service/src/pkg/secure"
	"banner-display-service/src/pkg/validator"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
)

func Run(configPath string) {

	// Config 
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Logger
	SetLogrus(cfg.Log.Level)

	// Database
	log.Info("Initializing postgres")
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.MaxPoolSize))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Repositories
	log.Info("Initializing repositories")
	repositories := repositories.NewRepositories(pg)

	// Services
	log.Info("Initializing services")
	deps := services.ServicesDependencies{
		Repos:     repositories,
		APISecure: secure.NewSecure(cfg.Salt),
	}
	services := services.NewServices(deps)

	// Echo
	log.Info("Initializing Echo")
	handler := echo.New()
	handler.Validator = validator.NewCustomValidator()
	v1.NewRouter(handler, services)

	// HTTP server
	log.Info("Starting http server")
	log.Debugf("Server port: %s", cfg.HTTP.Port)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
	
	// Waiting signal
	log.Info("Configuring graceful shutdown...")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	
	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Graceful shutdown
	log.Info("Shutting down...")
	err = httpServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}