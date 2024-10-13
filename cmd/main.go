package main

import (
	"context"
	"dietsense/internal/api"
	"dietsense/internal/repositories"
	"dietsense/internal/repositories/postgres"
	"dietsense/internal/repositories/sqlite"
	"dietsense/internal/services"
	"dietsense/pkg/config"
	"dietsense/pkg/logging"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set up configuration
	config.Setup()

	// Initialize logging
	logging.Setup()
	logger := logging.Log // Use the global logger instance from logging package

	// Initialize the database based on configuration
	var db repositories.Database
	var err error
	if config.Config.DatabaseType == "postgres" {
		dsn := config.Config.DatabaseURL
		db, err = postgres.NewPostgresDB(dsn)
	} else if config.Config.DatabaseType == "sqlite" {
		dsn := config.Config.DatabaseURL
		db, err = sqlite.NewSQLiteDB(dsn)
	} else {
		logger.Fatalf("Unsupported database type: %s", config.Config.DatabaseType)
	}
	if err != nil {
		logger.Fatalf("Failed to connect to database: %s", err)
	}

	// Set up the service factory
	factory := services.NewServiceFactory(&config.Config)

	// Set up the Gin router with logging middleware
	router := gin.New()                   // Creates a router without any middleware by default
	router.Use(gin.Recovery())            // Adds built-in recovery middleware
	router.Use(logging.GinLogger(logger)) // Use custom Logrus-based logger middleware

	// Set up API routes, pass the db instance to handlers or services
	api.SetupRoutes(router, factory, db)

	// Create the HTTP server
	server := &http.Server{
		Addr:    config.Config.ServerAddress, // Use server address from config
		Handler: router,
	}
	logger.Info("Server Address:" + config.Config.ServerAddress)
	// Start server in a goroutine
	go func() {
		logger.Infof("Starting server on %s", config.Config.ServerAddress)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Server failed: %s", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	// Shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: %s", err)
	}

	logger.Info("Server exiting")
}
