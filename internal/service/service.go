package service

import (
	"smrtmart-go-postgresql/internal/config"
	"smrtmart-go-postgresql/internal/repository"
)

type Services struct {
	User     UserService
	Vendor   VendorService
	Product  ProductService
	Order    OrderService
	Cart     CartService
	Category CategoryService
	Review   ReviewService
	Payment  PaymentService
	Auth     AuthService
	Upload   UploadService
}

func NewServices(repos *repository.Repositories, cfg *config.Config) *Services {
	return &Services{
		User:     NewUserService(repos.User),
		Vendor:   NewVendorService(repos.Vendor),
		Product:  NewProductService(repos.Product),
		Order:    NewOrderService(repos.Order, repos.Product),
		Cart:     NewCartService(repos.Cart, repos.Product),
		Category: NewCategoryService(repos.Category),
		Review:   NewReviewService(repos.Review),
		Payment:  NewPaymentService(cfg.Stripe),
		Auth:     NewAuthService(repos.User, cfg.JWT),
		Upload:   NewUploadService(cfg.Upload),
	}
}