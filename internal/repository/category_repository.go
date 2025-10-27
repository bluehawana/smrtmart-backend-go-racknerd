package repository

import (
	"database/sql"
	"smrtmart-go-postgresql/internal/models"
)

type CategoryRepository interface {
	GetAll() ([]models.Category, error)
	GetByID(id string) (*models.Category, error)
	Create(category *models.Category) error
	Update(category *models.Category) error
	Delete(id string) error
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetAll() ([]models.Category, error) {
	query := `
		SELECT id, name, slug, description, image, parent_id, sort_order, is_active, created_at, updated_at
		FROM categories
		WHERE is_active = true
		ORDER BY sort_order ASC, name ASC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Slug,
			&c.Description,
			&c.Image,
			&c.ParentID,
			&c.SortOrder,
			&c.IsActive,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, rows.Err()
}

func (r *categoryRepository) GetByID(id string) (*models.Category, error) {
	query := `
		SELECT id, name, slug, description, image, parent_id, sort_order, is_active, created_at, updated_at
		FROM categories
		WHERE id = ? AND is_active = true
	`

	var c models.Category
	err := r.db.QueryRow(query, id).Scan(
		&c.ID,
		&c.Name,
		&c.Slug,
		&c.Description,
		&c.Image,
		&c.ParentID,
		&c.SortOrder,
		&c.IsActive,
		&c.CreatedAt,
		&c.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *categoryRepository) Create(category *models.Category) error {
	// TODO: Implement create
	return nil
}

func (r *categoryRepository) Update(category *models.Category) error {
	// TODO: Implement update
	return nil
}

func (r *categoryRepository) Delete(id string) error {
	// TODO: Implement delete
	return nil
}
