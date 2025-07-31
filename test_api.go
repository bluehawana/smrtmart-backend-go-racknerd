package main

import (
	"log"
	"net/http"
	"strconv"

	"smrtmart-backend/internal/api"
	"smrtmart-backend/internal/config"
	"smrtmart-backend/internal/database"
	"smrtmart-backend/internal/repository"
	"smrtmart-backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := database.Initialize(cfg.Database)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Initialize simple repository
	repo := repository.NewSimpleProductRepository(db)

	// Setup Gin
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "SmartMart API",
			"version": "1.0.0",
		})
	})

	// Products endpoint
	router.GET("/api/v1/products", func(c *gin.Context) {
		filters := repository.ProductFilters{
			Status: c.DefaultQuery("status", "active"),
		}

		// Parse limit
		if limitStr := c.Query("limit"); limitStr != "" {
			if limit, err := strconv.Atoi(limitStr); err == nil && limit > 0 {
				filters.Limit = limit
			}
		}

		products, total, err := repo.GetAll(filters)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to get products",
				"error":   err.Error(),
			})
			return
		}

		if filters.Limit <= 0 {
			filters.Limit = 20
		}
		if filters.Page <= 0 {
			filters.Page = 1
		}
		totalPages := (total + filters.Limit - 1) / filters.Limit

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Products retrieved successfully",
			"data": gin.H{
				"data": products,
				"pagination": gin.H{
					"page":        filters.Page,
					"limit":       filters.Limit,
					"total":       total,
					"total_pages": totalPages,
				},
			},
		})
	})

	// Featured products endpoint
	router.GET("/api/v1/products/featured", func(c *gin.Context) {
		limit := 10
		if limitStr := c.Query("limit"); limitStr != "" {
			if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 {
				limit = parsedLimit
			}
		}

		products, err := repo.GetFeatured(limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to get featured products",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Featured products retrieved successfully",
			"data":    products,
		})
	})

	// Payment service and handler
	paymentService := service.NewPaymentService(cfg.Stripe)
	paymentHandler := api.NewPaymentHandler(paymentService)

	// Payment endpoints
	router.POST("/api/v1/orders/checkout", paymentHandler.CreateCheckoutSession)
	router.POST("/api/v1/webhooks/stripe", paymentHandler.StripeWebhook)

	// Static file serving
	router.Static("/uploads", "./uploads")

	log.Println("🚀 SmartMart API server starting on port 8080")
	log.Println("💳 Stripe integration enabled")
	log.Println("📚 Endpoints available:")
	log.Println("   GET  /api/v1/products")
	log.Println("   GET  /api/v1/products/featured")
	log.Println("   POST /api/v1/orders/checkout")
	log.Println("   POST /api/v1/webhooks/stripe")
	log.Fatal(router.Run(":8080"))
}