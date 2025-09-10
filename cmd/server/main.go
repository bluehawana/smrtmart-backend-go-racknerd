package main

import (
	"embed"
	"log"
	"os"

	"smrtmart-go-postgresql/internal/api"
	"smrtmart-go-postgresql/internal/config"
	"smrtmart-go-postgresql/internal/database"
	"smrtmart-go-postgresql/internal/middleware"
	"smrtmart-go-postgresql/internal/repository"
	"smrtmart-go-postgresql/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

// @title SmrtMart API
// @version 1.0
// @description Professional ecommerce platform API for SME digitalization
// @termsOfService https://smrtmart.com/terms

// @contact.name SmrtMart Support
// @contact.url https://smrtmart.com/support
// @contact.email support@smrtmart.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host api.smrtmart.com
// @BasePath /api/v1
// @schemes https http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := database.Initialize(cfg.Database)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Run migrations
	if err := database.RunMigrations(cfg.Database); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Initialize repositories
	repos := repository.NewRepositories(db)

	// Initialize services
	services := service.NewServices(repos, cfg)

	// Initialize Gin router
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	
	router := gin.New()
	
	// Add middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORS(cfg.Server.CORSOrigins))
	router.Use(middleware.SecurityHeaders())
	router.Use(middleware.RateLimit())

	// Setup routes
	api.SetupRoutes(router, services, cfg)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 SmrtMart API server starting on port %s", port)
	log.Printf("📚 API Documentation: http://localhost:%s/swagger/index.html", port)
	
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}