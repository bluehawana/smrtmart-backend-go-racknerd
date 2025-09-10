package api

import (
	"net/http"

	"smrtmart-go-postgresql/internal/config"
	"smrtmart-go-postgresql/internal/middleware"
	"smrtmart-go-postgresql/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(router *gin.Engine, services *service.Services, cfg *config.Config) {
	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "SmrtMart API",
			"version": "1.0.0",
		})
	})
	
	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API v1 routes
	v1 := router.Group("/api/v1")
	
	// Health check endpoint for API v1
	v1.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "SmrtMart API v1",
			"version": "1.0.0",
		})
	})
	{
		// Public routes
		public := v1.Group("/")
		{
			// Products
			products := public.Group("/products")
			{
				productHandler := NewProductHandler(services.Product)
				products.GET("", productHandler.GetProducts)
				products.GET("/:id", productHandler.GetProduct)
				products.GET("/search", productHandler.SearchProducts)
				products.GET("/featured", productHandler.GetFeaturedProducts)
			}

			// Categories
			categories := public.Group("/categories")
			{
				categoryHandler := NewCategoryHandler(services.Category)
				categories.GET("", categoryHandler.GetCategories)
				categories.GET("/:id", categoryHandler.GetCategory)
			}

			// Authentication
			auth := public.Group("/auth")
			{
				authHandler := NewAuthHandler(services.Auth)
				auth.POST("/register", authHandler.Register)
				auth.POST("/login", authHandler.Login)
				auth.POST("/refresh", authHandler.RefreshToken)
				auth.POST("/forgot-password", authHandler.ForgotPassword)
				auth.POST("/reset-password", authHandler.ResetPassword)
			}

			// Cart (guest and authenticated)
			cart := public.Group("/cart")
			{
				cartHandler := NewCartHandler(services.Cart)
				cart.Use(middleware.OptionalJWTAuth())
				cart.GET("", cartHandler.GetCart)
				cart.POST("/items", cartHandler.AddItem)
				cart.PUT("/items/:id", cartHandler.UpdateItem)
				cart.DELETE("/items/:id", cartHandler.RemoveItem)
				cart.DELETE("", cartHandler.ClearCart)
				cart.POST("/clear", cartHandler.ClearCart)  // Additional endpoint for frontend compatibility
			}

			// Orders (checkout)
			orders := public.Group("/orders")
			{
				paymentHandler := NewPaymentHandler(services.Payment)
				orders.POST("/checkout", paymentHandler.CreateCheckoutSession)
			}

			// Payment webhooks
			webhooks := public.Group("/webhooks")
			{
				paymentHandler := NewPaymentHandler(services.Payment)
				webhooks.POST("/stripe", paymentHandler.StripeWebhook)
			}
		}

		// Protected routes (require authentication)
		protected := v1.Group("/")
		protected.Use(middleware.JWTAuth())
		{
			// User profile
			users := protected.Group("/users")
			{
				userHandler := NewUserHandler(services.User)
				users.GET("/profile", userHandler.GetProfile)
				users.PUT("/profile", userHandler.UpdateProfile)
				users.POST("/change-password", userHandler.ChangePassword)
			}

			// User orders
			orders := protected.Group("/orders")
			{
				orderHandler := NewOrderHandler(services.Order)
				orders.GET("", orderHandler.GetUserOrders)
				orders.GET("/:id", orderHandler.GetOrder)
				orders.POST("/:id/cancel", orderHandler.CancelOrder)
			}

			// Reviews
			reviews := protected.Group("/reviews")
			{
				reviewHandler := NewReviewHandler(services.Review)
				reviews.POST("", reviewHandler.CreateReview)
				reviews.PUT("/:id", reviewHandler.UpdateReview)
				reviews.DELETE("/:id", reviewHandler.DeleteReview)
			}
		}

		// Vendor routes
		vendor := v1.Group("/vendor")
		vendor.Use(middleware.JWTAuth(), middleware.RequireVendor())
		{
			// Vendor profile
			vendor.GET("/profile", func(c *gin.Context) {
				// TODO: Implement vendor profile handler
				c.JSON(http.StatusOK, gin.H{"message": "Vendor profile endpoint"})
			})

			// Vendor products
			products := vendor.Group("/products")
			{
				productHandler := NewProductHandler(services.Product)
				products.GET("", productHandler.GetVendorProducts)
				products.POST("", productHandler.CreateProduct)
				products.PUT("/:id", productHandler.UpdateProduct)
				products.DELETE("/:id", productHandler.DeleteProduct)
				products.PATCH("/:id/stock", productHandler.UpdateProductStock)
			}

			// Vendor orders
			orders := vendor.Group("/orders")
			{
				orderHandler := NewOrderHandler(services.Order)
				orders.GET("", orderHandler.GetVendorOrders)
				orders.PUT("/:id/status", orderHandler.UpdateOrderStatus)
			}
		}

		// Admin routes
		admin := v1.Group("/admin")
		admin.Use(middleware.JWTAuth(), middleware.RequireAdmin())
		{
			// User management
			users := admin.Group("/users")
			{
				userHandler := NewUserHandler(services.User)
				users.GET("", userHandler.GetUsers)
				users.GET("/:id", userHandler.GetUser)
				users.PUT("/:id/status", userHandler.UpdateUserStatus)
				users.DELETE("/:id", userHandler.DeleteUser)
			}

			// Vendor management
			vendors := admin.Group("/vendors")
			{
				vendorHandler := NewVendorHandler(services.Vendor)
				vendors.GET("", vendorHandler.GetVendors)
				vendors.GET("/:id", vendorHandler.GetVendor)
				vendors.PUT("/:id/status", vendorHandler.UpdateVendorStatus)
				vendors.POST("/:id/verify", vendorHandler.VerifyVendor)
			}

			// Product management
			products := admin.Group("/products")
			{
				productHandler := NewProductHandler(services.Product)
				products.GET("", productHandler.GetAllProducts)
				products.PUT("/:id/featured", productHandler.ToggleFeatured)
				products.PUT("/:id/status", productHandler.UpdateProductStatus)
			}

			// Category management
			categories := admin.Group("/categories")
			{
				categoryHandler := NewCategoryHandler(services.Category)
				categories.POST("", categoryHandler.CreateCategory)
				categories.PUT("/:id", categoryHandler.UpdateCategory)
				categories.DELETE("/:id", categoryHandler.DeleteCategory)
			}

			// Order management
			orders := admin.Group("/orders")
			{
				orderHandler := NewOrderHandler(services.Order)
				orders.GET("", orderHandler.GetAllOrders)
				orders.PUT("/:id/status", orderHandler.UpdateOrderStatus)
			}

			// Database migration management
			migrationHandler := NewMigrationHandler(cfg)
			admin.POST("/migrate", migrationHandler.RunMigrations)
			admin.GET("/migrate/status", migrationHandler.GetMigrationStatus)
		}

		// File upload routes
		upload := v1.Group("/upload")
		upload.Use(middleware.JWTAuth())
		{
			uploadHandler := NewUploadHandler(services.Upload)
			upload.POST("/image", uploadHandler.UploadImage)
			upload.POST("/images", uploadHandler.UploadMultipleImages)
		}

		// Static file serving
		router.Static("/uploads", "./uploads")
	}
}