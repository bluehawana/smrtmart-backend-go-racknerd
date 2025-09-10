package api

import (
	"net/http"

	"smrtmart-go-postgresql/internal/config"
	"smrtmart-go-postgresql/internal/database"

	"github.com/gin-gonic/gin"
)

// MigrationHandler handles database migration endpoints
type MigrationHandler struct {
	config *config.Config
}

// NewMigrationHandler creates a new migration handler
func NewMigrationHandler(cfg *config.Config) *MigrationHandler {
	return &MigrationHandler{
		config: cfg,
	}
}

// RunMigrations runs database migrations up
// @Summary Run database migrations
// @Description Run all pending database migrations
// @Tags admin
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /admin/migrate [post]
func (h *MigrationHandler) RunMigrations(c *gin.Context) {
	// Run migrations
	if err := database.RunMigrations(h.config.Database); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Migration failed",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Database migrations completed successfully",
		"status":  "success",
	})
}

// GetMigrationStatus gets the current migration status
// @Summary Get migration status
// @Description Get the current database migration status
// @Tags admin
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /admin/migrate/status [get]
func (h *MigrationHandler) GetMigrationStatus(c *gin.Context) {
	// This would require implementing migration status check
	// For now, return a simple status
	c.JSON(http.StatusOK, gin.H{
		"message": "Migration status endpoint - implement migration version checking",
		"status":  "available",
	})
}