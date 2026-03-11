package main

import (
	"kms_golang/config"
	"kms_golang/database"
	"kms_golang/middleware"
	"kms_golang/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration from .env
	config.Load()

	// Set Gin mode
	if config.AppConfig.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Connect to database
	database.Connect()

	// Create Gin engine
	r := gin.Default()

	// Global middleware
	r.Use(middleware.CORS())

	// Register all routes
	routes.RegisterRoutes(r)

	// Start server
	addr := ":" + config.AppConfig.AppPort
	log.Printf("Starting KMS server on %s (env: %s)", addr, config.AppConfig.AppEnv)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
