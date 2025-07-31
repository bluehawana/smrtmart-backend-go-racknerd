package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"smrtmart-backend/internal/models"

	"github.com/lib/pq"
)

type SimpleProductRepository interface {
	GetAll(filters ProductFilters) ([]*models.Product, int, error)
	GetFeatured(limit int) ([]*models.Product, error)
}

type simpleProductRepository struct {
	db *sql.DB
}

func NewSimpleProductRepository(db *sql.DB) SimpleProductRepository {
	return &simpleProductRepository{db: db}
}

func (r *simpleProductRepository) GetAll(filters ProductFilters) ([]*models.Product, int, error) {
	whereClause, args := r.buildWhereClause(filters)
	
	// Count query
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM products %s", whereClause)
	var total int
	err := r.db.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Main query - simplified without JSON fields for now
	orderClause := r.buildOrderClause(filters)
	limitClause := r.buildLimitClause(filters)
	
	query := fmt.Sprintf(`
		SELECT id, vendor_id, name, description, price, compare_price, sku,
			category, tags, images, stock, status, featured, weight,
			created_at, updated_at
		FROM products %s %s %s`, whereClause, orderClause, limitClause)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		product := &models.Product{}
		err := rows.Scan(
			&product.ID, &product.VendorID, &product.Name, &product.Description,
			&product.Price, &product.ComparePrice, &product.SKU, &product.Category,
			pq.Array(&product.Tags), pq.Array(&product.Images), &product.Stock,
			&product.Status, &product.Featured, &product.Weight,
			&product.CreatedAt, &product.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		products = append(products, product)
	}

	return products, total, nil
}

func (r *simpleProductRepository) GetFeatured(limit int) ([]*models.Product, error) {
	query := `
		SELECT id, vendor_id, name, description, price, compare_price, sku,
			category, tags, images, stock, status, featured, weight,
			created_at, updated_at
		FROM products 
		WHERE featured = true AND status = 'active' AND stock > 0
		ORDER BY created_at DESC
		LIMIT $1`

	rows, err := r.db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		product := &models.Product{}
		err := rows.Scan(
			&product.ID, &product.VendorID, &product.Name, &product.Description,
			&product.Price, &product.ComparePrice, &product.SKU, &product.Category,
			pq.Array(&product.Tags), pq.Array(&product.Images), &product.Stock,
			&product.Status, &product.Featured, &product.Weight,
			&product.CreatedAt, &product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *simpleProductRepository) buildWhereClause(filters ProductFilters) (string, []interface{}) {
	var conditions []string
	var args []interface{}
	argIndex := 1

	if filters.Category != "" {
		conditions = append(conditions, fmt.Sprintf("category = $%d", argIndex))
		args = append(args, filters.Category)
		argIndex++
	}

	if filters.Status != "" {
		conditions = append(conditions, fmt.Sprintf("status = $%d", argIndex))
		args = append(args, filters.Status)
		argIndex++
	}

	if filters.Featured != nil {
		conditions = append(conditions, fmt.Sprintf("featured = $%d", argIndex))
		args = append(args, *filters.Featured)
		argIndex++
	}

	if filters.MinPrice != nil {
		conditions = append(conditions, fmt.Sprintf("price >= $%d", argIndex))
		args = append(args, *filters.MinPrice)
		argIndex++
	}

	if filters.MaxPrice != nil {
		conditions = append(conditions, fmt.Sprintf("price <= $%d", argIndex))
		args = append(args, *filters.MaxPrice)
		argIndex++
	}

	if len(conditions) == 0 {
		return "", args
	}

	return "WHERE " + strings.Join(conditions, " AND "), args
}

func (r *simpleProductRepository) buildOrderClause(filters ProductFilters) string {
	if filters.SortBy == "" {
		return "ORDER BY created_at DESC"
	}

	sortDir := "ASC"
	if filters.SortDir == "desc" {
		sortDir = "DESC"
	}

	switch filters.SortBy {
	case "name":
		return fmt.Sprintf("ORDER BY name %s", sortDir)
	case "price":
		return fmt.Sprintf("ORDER BY price %s", sortDir)
	case "created_at":
		return fmt.Sprintf("ORDER BY created_at %s", sortDir)
	case "updated_at":
		return fmt.Sprintf("ORDER BY updated_at %s", sortDir)
	default:
		return "ORDER BY created_at DESC"
	}
}

func (r *simpleProductRepository) buildLimitClause(filters ProductFilters) string {
	if filters.Limit <= 0 {
		filters.Limit = 20
	}
	if filters.Page <= 0 {
		filters.Page = 1
	}

	offset := (filters.Page - 1) * filters.Limit
	return fmt.Sprintf("LIMIT %d OFFSET %d", filters.Limit, offset)
}