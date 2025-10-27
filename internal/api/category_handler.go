package api

import (
	"net/http"

	"smrtmart-go-postgresql/internal/models"
	"smrtmart-go-postgresql/internal/service"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service service.CategoryService
}

func NewCategoryHandler(service service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// GetCategories godoc
// @Summary Get all categories
// @Description Get a list of all active categories
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /categories [get]
func (h *CategoryHandler) GetCategories(c *gin.Context) {
	categories, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to get categories",
			Error: &models.APIError{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Categories retrieved successfully",
		Data:    categories,
	})
}

// GetCategory godoc
// @Summary Get a category by ID
// @Description Get details of a specific category
// @Tags categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} models.APIResponse
// @Failure 404 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /categories/{id} [get]
func (h *CategoryHandler) GetCategory(c *gin.Context) {
	id := c.Param("id")

	category, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.APIResponse{
			Success: false,
			Message: "Category not found",
			Error: &models.APIError{
				Code:    "NOT_FOUND",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Category retrieved successfully",
		Data:    category,
	})
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Create category endpoint - TODO"})
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update category endpoint - TODO"})
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete category endpoint - TODO"})
}
