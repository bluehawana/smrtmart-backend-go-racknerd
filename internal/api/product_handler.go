package api

import (
	"net/http"
	"strconv"
	"regexp"

	"smrtmart-go-postgresql/internal/models"
	"smrtmart-go-postgresql/internal/repository"
	"smrtmart-go-postgresql/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// GetProducts godoc
// @Summary Get products with filtering and pagination
// @Description Get a list of products with optional filtering by category, price range, etc.
// @Tags products
// @Accept json
// @Produce json
// @Param category query string false "Filter by category"
// @Param status query string false "Filter by status" Enums(active, inactive, draft)
// @Param featured query bool false "Filter by featured status"
// @Param min_price query number false "Minimum price filter"
// @Param max_price query number false "Maximum price filter"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(20)
// @Param sort_by query string false "Sort by field" Enums(name, price, created_at)
// @Param sort_dir query string false "Sort direction" Enums(asc, desc) default(desc)
// @Success 200 {object} models.PaginatedResponse
// @Failure 400 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /products [get]
func (h *ProductHandler) GetProducts(c *gin.Context) {
	filters := repository.ProductFilters{
		Category: c.Query("category"),
		Status:   c.Query("status"),
		SortBy:   c.Query("sort_by"),
		SortDir:  c.Query("sort_dir"),
	}

	// Parse featured filter
	if featuredStr := c.Query("featured"); featuredStr != "" {
		if featured, err := strconv.ParseBool(featuredStr); err == nil {
			filters.Featured = &featured
		}
	}

	// Parse price filters
	if minPriceStr := c.Query("min_price"); minPriceStr != "" {
		if minPrice, err := strconv.ParseFloat(minPriceStr, 64); err == nil {
			filters.MinPrice = &minPrice
		}
	}
	if maxPriceStr := c.Query("max_price"); maxPriceStr != "" {
		if maxPrice, err := strconv.ParseFloat(maxPriceStr, 64); err == nil {
			filters.MaxPrice = &maxPrice
		}
	}

	// Parse pagination
	if pageStr := c.Query("page"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil && page > 0 {
			filters.Page = page
		}
	}
	if limitStr := c.Query("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil && limit > 0 {
			filters.Limit = limit
		}
	}

	result, err := h.service.GetProducts(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to get products",
			Error: &models.APIError{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Products retrieved successfully",
		Data:    result,
	})
}

// GetProduct godoc
// @Summary Get a product by ID
// @Description Get detailed information about a specific product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.APIResponse{data=models.Product}
// @Failure 400 {object} models.APIResponse
// @Failure 404 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /products/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	idStr := c.Param("id")
	
	// Check if the ID is numeric (1-50) or UUID format
	var product *models.Product
	var err error
	
	// Try parsing as numeric ID first (simpler and more user-friendly)
	if numericID, numErr := strconv.Atoi(idStr); numErr == nil && numericID >= 1 && numericID <= 50 {
		// It's a valid numeric ID, get product by numeric_id
		product, err = h.service.GetProductByNumericID(numericID)
	} else {
		// Try parsing as UUID
		uuidRegex := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`)
		if !uuidRegex.MatchString(idStr) {
			c.JSON(http.StatusBadRequest, models.APIResponse{
				Success: false,
				Message: "Invalid product ID format",
				Error: &models.APIError{
					Code:    "INVALID_ID",
					Message: "Product ID must be a valid UUID or numeric ID (1-50)",
				},
			})
			return
		}
		
		id, parseErr := uuid.Parse(idStr)
		if parseErr != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse{
				Success: false,
				Message: "Invalid UUID format",
				Error: &models.APIError{
					Code:    "INVALID_UUID",
					Message: "Product UUID is malformed",
				},
			})
			return
		}
		
		product, err = h.service.GetProduct(id)
	}
	if err != nil {
		if err.Error() == "product not found" {
			c.JSON(http.StatusNotFound, models.APIResponse{
				Success: false,
				Message: "Product not found",
				Error: &models.APIError{
					Code:    "NOT_FOUND",
					Message: "Product with the specified ID does not exist",
				},
			})
			return
		}

		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to get product",
			Error: &models.APIError{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Product retrieved successfully",
		Data:    product,
	})
}

// SearchProducts godoc
// @Summary Search products
// @Description Search products by name and description
// @Tags products
// @Accept json
// @Produce json
// @Param q query string true "Search query"
// @Param category query string false "Filter by category"
// @Param min_price query number false "Minimum price filter"
// @Param max_price query number false "Maximum price filter"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(20)
// @Success 200 {object} models.PaginatedResponse
// @Failure 400 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /products/search [get]
func (h *ProductHandler) SearchProducts(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Search query is required",
			Error: &models.APIError{
				Code:    "MISSING_QUERY",
				Message: "Search query parameter 'q' is required",
			},
		})
		return
	}

	filters := repository.ProductFilters{
		Category: c.Query("category"),
		Status:   "active", // Only search active products
	}

	// Parse price filters
	if minPriceStr := c.Query("min_price"); minPriceStr != "" {
		if minPrice, err := strconv.ParseFloat(minPriceStr, 64); err == nil {
			filters.MinPrice = &minPrice
		}
	}
	if maxPriceStr := c.Query("max_price"); maxPriceStr != "" {
		if maxPrice, err := strconv.ParseFloat(maxPriceStr, 64); err == nil {
			filters.MaxPrice = &maxPrice
		}
	}

	// Parse pagination
	if pageStr := c.Query("page"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil && page > 0 {
			filters.Page = page
		}
	}
	if limitStr := c.Query("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil && limit > 0 {
			filters.Limit = limit
		}
	}

	result, err := h.service.SearchProducts(query, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to search products",
			Error: &models.APIError{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Search completed successfully",
		Data:    result,
	})
}

// GetFeaturedProducts godoc
// @Summary Get featured products
// @Description Get a list of featured products
// @Tags products
// @Accept json
// @Produce json
// @Param limit query int false "Number of products to return" default(10)
// @Success 200 {object} models.APIResponse{data=[]models.Product}
// @Failure 500 {object} models.APIResponse
// @Router /products/featured [get]
func (h *ProductHandler) GetFeaturedProducts(c *gin.Context) {
	limit := 10
	if limitStr := c.Query("limit"); limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	products, err := h.service.GetFeaturedProducts(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to get featured products",
			Error: &models.APIError{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Featured products retrieved successfully",
		Data:    products,
	})
}

// CreateProduct godoc
// @Summary Create a new product (Vendor only)
// @Description Create a new product for the authenticated vendor
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param product body models.Product true "Product data"
// @Success 201 {object} models.APIResponse{data=models.Product}
// @Failure 400 {object} models.APIResponse
// @Failure 401 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /vendor/products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid request data",
			Error: &models.APIError{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		})
		return
	}

	// TODO: Get vendor ID from JWT token
	// For now, we'll use a placeholder
	product.VendorID = uuid.New()

	if err := h.service.CreateProduct(&product); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Failed to create product",
			Error: &models.APIError{
				Code:    "CREATION_FAILED",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusCreated, models.APIResponse{
		Success: true,
		Message: "Product created successfully",
		Data:    product,
	})
}

// UpdateProduct godoc
// @Summary Update a product (Vendor only)
// @Description Update an existing product owned by the authenticated vendor
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Product ID"
// @Param product body models.Product true "Updated product data"
// @Success 200 {object} models.APIResponse{data=models.Product}
// @Failure 400 {object} models.APIResponse
// @Failure 401 {object} models.APIResponse
// @Failure 404 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /vendor/products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid product ID",
			Error: &models.APIError{
				Code:    "INVALID_ID",
				Message: "Product ID must be a valid UUID",
			},
		})
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid request data",
			Error: &models.APIError{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		})
		return
	}

	product.ID = id

	if err := h.service.UpdateProduct(&product); err != nil {
		if err.Error() == "product not found" {
			c.JSON(http.StatusNotFound, models.APIResponse{
				Success: false,
				Message: "Product not found",
				Error: &models.APIError{
					Code:    "NOT_FOUND",
					Message: "Product with the specified ID does not exist",
				},
			})
			return
		}

		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Failed to update product",
			Error: &models.APIError{
				Code:    "UPDATE_FAILED",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Product updated successfully",
		Data:    product,
	})
}

// DeleteProduct godoc
// @Summary Delete a product (Vendor only)
// @Description Delete an existing product owned by the authenticated vendor
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Product ID"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 401 {object} models.APIResponse
// @Failure 404 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /vendor/products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid product ID",
			Error: &models.APIError{
				Code:    "INVALID_ID",
				Message: "Product ID must be a valid UUID",
			},
		})
		return
	}

	if err := h.service.DeleteProduct(id); err != nil {
		if err.Error() == "product not found" {
			c.JSON(http.StatusNotFound, models.APIResponse{
				Success: false,
				Message: "Product not found",
				Error: &models.APIError{
					Code:    "NOT_FOUND",
					Message: "Product with the specified ID does not exist",
				},
			})
			return
		}

		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to delete product",
			Error: &models.APIError{
				Code:    "DELETION_FAILED",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Product deleted successfully",
	})
}

// GetVendorProducts godoc
// @Summary Get vendor's products (Vendor only)
// @Description Get all products owned by the authenticated vendor
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param status query string false "Filter by status"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(20)
// @Success 200 {object} models.PaginatedResponse
// @Failure 401 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /vendor/products [get]
func (h *ProductHandler) GetVendorProducts(c *gin.Context) {
	// TODO: Get vendor ID from JWT token
	vendorID := uuid.New() // Placeholder

	filters := repository.ProductFilters{
		Status: c.Query("status"),
	}

	// Parse pagination
	if pageStr := c.Query("page"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil && page > 0 {
			filters.Page = page
		}
	}
	if limitStr := c.Query("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil && limit > 0 {
			filters.Limit = limit
		}
	}

	result, err := h.service.GetVendorProducts(vendorID, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to get vendor products",
			Error: &models.APIError{
				Code:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Vendor products retrieved successfully",
		Data:    result,
	})
}

// UpdateProductStock godoc
// @Summary Update product stock (Vendor only)
// @Description Update the stock quantity of a product
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Product ID"
// @Param stock body object{stock=int} true "Stock data"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 401 {object} models.APIResponse
// @Failure 404 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /vendor/products/{id}/stock [patch]
func (h *ProductHandler) UpdateProductStock(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid product ID",
			Error: &models.APIError{
				Code:    "INVALID_ID",
				Message: "Product ID must be a valid UUID",
			},
		})
		return
	}

	var req struct {
		Stock int `json:"stock" binding:"required,min=0"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid request data",
			Error: &models.APIError{
				Code:    "INVALID_REQUEST",
				Message: err.Error(),
			},
		})
		return
	}

	if err := h.service.UpdateProductStock(id, req.Stock); err != nil {
		if err.Error() == "product not found" {
			c.JSON(http.StatusNotFound, models.APIResponse{
				Success: false,
				Message: "Product not found",
				Error: &models.APIError{
					Code:    "NOT_FOUND",
					Message: "Product with the specified ID does not exist",
				},
			})
			return
		}

		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to update product stock",
			Error: &models.APIError{
				Code:    "UPDATE_FAILED",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Product stock updated successfully",
	})
}

// Admin-only handlers (placeholders)
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Admin get all products endpoint"})
}

func (h *ProductHandler) ToggleFeatured(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Admin toggle featured endpoint"})
}

func (h *ProductHandler) UpdateProductStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Admin update product status endpoint"})
}