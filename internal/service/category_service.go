package service

import (
	"smrtmart-go-postgresql/internal/models"
	"smrtmart-go-postgresql/internal/repository"
)

type CategoryService interface {
	GetAll() ([]models.Category, error)
	GetByID(id string) (*models.Category, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) GetAll() ([]models.Category, error) {
	return s.repo.GetAll()
}

func (s *categoryService) GetByID(id string) (*models.Category, error) {
	return s.repo.GetByID(id)
}
