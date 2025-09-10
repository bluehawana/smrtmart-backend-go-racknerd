package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// CORS middleware
func CORS(allowedOrigins []string) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		
		// Check if origin is allowed
		allowedOrigin := ""
		for _, allowed := range allowedOrigins {
			if origin == allowed {
				allowedOrigin = origin
				break
			}
		}
		
		// If no specific origins provided, allow common production domains
		if len(allowedOrigins) == 0 {
			commonOrigins := []string{
				"https://www.smrtmart.com",
				"https://smrtmart.com", 
				"http://localhost:3000",
				"http://localhost:3001",
			}
			for _, allowed := range commonOrigins {
				if origin == allowed {
					allowedOrigin = origin
					break
				}
			}
		}
		
		// Set CORS headers
		if allowedOrigin != "" {
			c.Header("Access-Control-Allow-Origin", allowedOrigin)
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Requested-With")
		c.Header("Access-Control-Expose-Headers", "Content-Length")
		c.Header("Access-Control-Max-Age", "43200") // 12 hours
		
		// Debug info
		c.Header("X-Debug-Origin", origin)
		c.Header("X-Debug-Allowed-Origin", allowedOrigin)
		
		// Handle preflight OPTIONS request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		
		c.Next()
	})
}

// Security headers middleware
func SecurityHeaders() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Header("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'; img-src 'self' data: https:; font-src 'self' https:; connect-src 'self' https:; media-src 'self'; object-src 'none'; child-src 'none'; worker-src 'none'; frame-ancestors 'none'; form-action 'self'; base-uri 'self'; manifest-src 'self'")
		
		c.Next()
	})
}

// Rate limiting middleware
func RateLimit() gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Every(time.Minute), 100) // 100 requests per minute

	return gin.HandlerFunc(func(c *gin.Context) {
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"success": false,
				"message": "Rate limit exceeded",
				"error": gin.H{
					"code":    "RATE_LIMIT_EXCEEDED",
					"message": "Too many requests, please try again later",
				},
			})
			c.Abort()
			return
		}
		c.Next()
	})
}

// JWT Authentication middleware
func JWTAuth() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Authorization header required",
				"error": gin.H{
					"code":    "UNAUTHORIZED",
					"message": "Authorization header is required",
				},
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid authorization format",
				"error": gin.H{
					"code":    "INVALID_TOKEN_FORMAT",
					"message": "Authorization header must be in format: Bearer <token>",
				},
			})
			c.Abort()
			return
		}

		// TODO: Implement JWT token validation
		// For now, we'll skip validation and continue
		c.Next()
	})
}

// Optional JWT Authentication middleware (doesn't block if no token)
func OptionalJWTAuth() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString != authHeader {
				// TODO: Implement JWT token validation and set user context
			}
		}
		c.Next()
	})
}

// Admin role middleware
func RequireAdmin() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		// TODO: Check if user has admin role
		// For now, we'll allow all authenticated users
		c.Next()
	})
}

// Vendor role middleware
func RequireVendor() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		// TODO: Check if user has vendor role
		// For now, we'll allow all authenticated users
		c.Next()
	})
}