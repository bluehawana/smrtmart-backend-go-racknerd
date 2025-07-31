package service

import (
	"errors"

	"smrtmart-go-postgresql/internal/models"
	"smrtmart-go-postgresql/internal/repository"

	"github.com/google/uuid"
)

type ProductService interface {
	CreateProduct(product *models.Product) error
	GetProduct(id uuid.UUID) (*models.Product, error)
	GetProducts(filters repository.ProductFilters) (*models.PaginatedResponse, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uuid.UUID) error
	GetVendorProducts(vendorID uuid.UUID, filters repository.ProductFilters) (*models.PaginatedResponse, error)
	SearchProducts(query string, filters repository.ProductFilters) (*models.PaginatedResponse, error)
	GetFeaturedProducts(limit int) ([]*models.Product, error)
	UpdateProductStock(id uuid.UUID, stock int) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(product *models.Product) error {
	if product.Name == "" {
		return errors.New("product name is required")
	}
	if product.Price <= 0 {
		return errors.New("product price must be greater than 0")
	}
	if product.Category == "" {
		return errors.New("product category is required")
	}

	// Set default status if not provided
	if product.Status == "" {
		product.Status = models.ProductStatusDraft
	}

	// Initialize empty slices if nil
	if product.Tags == nil {
		product.Tags = []string{}
	}
	if product.Images == nil {
		product.Images = []string{}
	}

	return s.repo.Create(product)
}

func (s *productService) GetProduct(id uuid.UUID) (*models.Product, error) {
	if id == uuid.Nil {
		return nil, errors.New("invalid product ID")
	}

	product, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}

func (s *productService) GetProducts(filters repository.ProductFilters) (*models.PaginatedResponse, error) {
	// Set default pagination
	if filters.Limit <= 0 {
		filters.Limit = 20
	}
	if filters.Page <= 0 {
		filters.Page = 1
	}

	products, total, err := s.repo.GetAll(filters)
	if err != nil {
		return nil, err
	}

	totalPages := (total + filters.Limit - 1) / filters.Limit

	return &models.PaginatedResponse{
		Data: products,
		Pagination: models.Pagination{
			Page:       filters.Page,
			Limit:      filters.Limit,
			Total:      total,
			TotalPages: totalPages,
		},
	}, nil
}

func (s *productService) UpdateProduct(product *models.Product) error {
	if product.ID == uuid.Nil {
		return errors.New("invalid product ID")
	}
	if product.Name == "" {
		return errors.New("product name is required")
	}
	if product.Price <= 0 {
		return errors.New("product price must be greater than 0")
	}
	if product.Category == "" {
		return errors.New("product category is required")
	}

	// Check if product exists
	existing, err := s.repo.GetByID(product.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("product not found")
	}

	return s.repo.Update(product)
}

func (s *productService) DeleteProduct(id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("invalid product ID")
	}

	// Check if product exists
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("product not found")
	}

	return s.repo.Delete(id)
}

func (s *productService) GetVendorProducts(vendorID uuid.UUID, filters repository.ProductFilters) (*models.PaginatedResponse, error) {
	if vendorID == uuid.Nil {
		return nil, errors.New("invalid vendor ID")
	}

	// Set default pagination
	if filters.Limit <= 0 {
		filters.Limit = 20
	}
	if filters.Page <= 0 {
		filters.Page = 1
	}

	products, total, err := s.repo.GetByVendor(vendorID, filters)
	if err != nil {
		return nil, err
	}

	totalPages := (total + filters.Limit - 1) / filters.Limit

	return &models.PaginatedResponse{
		Data: products,
		Pagination: models.Pagination{
			Page:       filters.Page,
			Limit:      filters.Limit,
			Total:      total,
			TotalPages: totalPages,
		},
	}, nil
}

func (s *productService) SearchProducts(query string, filters repository.ProductFilters) (*models.PaginatedResponse, error) {
	if query == "" {
		return s.GetProducts(filters)
	}

	// Set default pagination
	if filters.Limit <= 0 {
		filters.Limit = 20
	}
	if filters.Page <= 0 {
		filters.Page = 1
	}

	products, total, err := s.repo.Search(query, filters)
	if err != nil {
		return nil, err
	}

	totalPages := (total + filters.Limit - 1) / filters.Limit

	return &models.PaginatedResponse{
		Data: products,
		Pagination: models.Pagination{
			Page:       filters.Page,
			Limit:      filters.Limit,
			Total:      total,
			TotalPages: totalPages,
		},
	}, nil
}

func (s *productService) GetFeaturedProducts(limit int) ([]*models.Product, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 50 {
		limit = 50
	}

	return s.repo.GetFeatured(limit)
}

func (s *productService) UpdateProductStock(id uuid.UUID, stock int) error {
	if id == uuid.Nil {
		return errors.New("invalid product ID")
	}
	if stock < 0 {
		return errors.New("stock cannot be negative")
	}

	// Check if product exists
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("product not found")
	}

	return s.repo.UpdateStock(id, stock)
}