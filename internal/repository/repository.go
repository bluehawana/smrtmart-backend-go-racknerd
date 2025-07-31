package repository

import (
	"database/sql"
)

type Repositories struct {
	User     UserRepository
	Vendor   VendorRepository
	Product  ProductRepository
	Order    OrderRepository
	Cart     CartRepository
	Category CategoryRepository
	Review   ReviewRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		User:     NewUserRepository(db),
		Vendor:   NewVendorRepository(db),
		Product:  NewProductRepository(db),
		Order:    NewOrderRepository(db),
		Cart:     NewCartRepository(db),
		Category: NewCategoryRepository(db),
		Review:   NewReviewRepository(db),
	}
}