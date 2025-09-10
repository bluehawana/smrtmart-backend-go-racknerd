# 🚀 SmrtMart Go Backend API

A professional ecommerce platform API built with **Go**, designed to help SMEs achieve digitalization through online selling, business management, and advertising solutions.

## 🚀 Features

### Core Ecommerce
- **Product Management**: Full CRUD operations with advanced filtering, search, and categorization
- **Order Processing**: Complete order lifecycle management with status tracking
- **Shopping Cart**: Session-based and user-based cart management
- **Payment Integration**: Stripe payment processing with webhook support
- **User Management**: Customer, vendor, and admin role-based access control

### SME Platform Features
- **Multi-Vendor Support**: Enable multiple businesses to sell on the platform
- **Vendor Dashboard**: Business management tools for SMEs
- **Advanced Analytics**: Sales reporting and business insights
- **Digital Marketing**: Built-in advertising and promotion tools
- **Professional APIs**: RESTful APIs with comprehensive documentation

### Technical Excellence
- **Scalable Architecture**: Clean architecture with repository pattern
- **Database**: PostgreSQL with optimized queries and indexing
- **Security**: JWT authentication, rate limiting, CORS, security headers
- **Performance**: Redis caching, connection pooling, optimized queries
- **Documentation**: Swagger/OpenAPI 3.0 documentation
- **Deployment**: Docker containerization with Docker Compose

## 🛠 Tech Stack

- **Language**: Go 1.21+
- **Framework**: Gin (High-performance HTTP router)
- **Database**: PostgreSQL 15 with optimized queries
- **Cache**: Redis 7 (ready for scaling)
- **Payment**: Stripe API with webhook support
- **Authentication**: JWT tokens with role-based access
- **Documentation**: Swagger/OpenAPI 3.0
- **Deployment**: Docker, Railway, Vercel
- **Migration**: golang-migrate with version control

## 📋 Prerequisites

- Go 1.21 or higher
- PostgreSQL 15
- Redis 7
- Docker & Docker Compose (for containerized deployment)
- Make (optional, for using Makefile commands)

## 🚀 Quick Start

### Option 1: Quick Start with Go (Recommended)

1. **Clone and setup**:
   ```bash
   git clone https://github.com/bluehawana/smrtmart-go-backend.git
   cd smrtmart-go-backend
   cp .env.example .env
   ```

2. **Configure environment**:
   Edit `.env` file with your settings (database, Stripe keys, etc.)

3. **Start the Go server**:
   ```bash
   go run test_api.go
   # or use the deployment script
   ./deploy.sh
   ```

4. **Verify deployment**:
   ```bash
   curl http://localhost:8080/health
   ```

5. **Access API documentation**:
   Open http://localhost:8080/swagger/index.html

### Option 2: Local Development

1. **Setup database**:
   ```bash
   # Start PostgreSQL and Redis locally
   # Create database: smrtmart_db
   ```

2. **Install dependencies**:
   ```bash
   make deps
   # or
   go mod download
   ```

3. **Run migrations**:
   ```bash
   make migrate-up
   ```

4. **Start the server**:
   ```bash
   make run
   # or
   go run ./cmd/server/main.go
   ```

## 📚 API Documentation

### Base URL
- Development: `http://localhost:8080/api/v1`
- Production: `https://api.smrtmart.com/api/v1`

### Authentication
```bash
# Register/Login to get JWT token
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password"}'

# Use token in subsequent requests
curl -H "Authorization: Bearer <your-jwt-token>" \
  http://localhost:8080/api/v1/users/profile
```

### Key Endpoints

#### Products
- `GET /api/v1/products` - List products with filtering
- `GET /api/v1/products/{id}` - Get product details
- `GET /api/v1/products/search?q=query` - Search products
- `GET /api/v1/products/featured` - Get featured products

#### Vendor Management
- `POST /api/v1/vendor/products` - Create product (vendor only)
- `GET /api/v1/vendor/products` - Get vendor's products
- `PUT /api/v1/vendor/products/{id}` - Update product
- `DELETE /api/v1/vendor/products/{id}` - Delete product

#### Orders & Cart
- `GET /api/v1/cart` - Get shopping cart
- `POST /api/v1/cart/items` - Add item to cart
- `POST /api/v1/orders/checkout` - Create order
- `GET /api/v1/orders` - Get user orders

### Complete API Documentation
Visit `/swagger/index.html` for interactive API documentation.

## 🏗 Architecture

```
smrtmart-go-backend/
├── cmd/server/          # Application entry point
├── internal/
│   ├── api/            # HTTP handlers and routes
│   ├── config/         # Configuration management
│   ├── database/       # Database connection and migrations
│   ├── middleware/     # HTTP middleware (auth, CORS, etc.)
│   ├── models/         # Data models and structs
│   ├── repository/     # Data access layer
│   └── service/        # Business logic layer
├── migrations/         # Database migrations
├── docs/              # Swagger documentation
├── uploads/           # File uploads directory
├── docker-compose.yml # Docker services configuration
├── Dockerfile         # Container build instructions
└── Makefile          # Development commands
```

## 🔧 Configuration

### Environment Variables

```bash
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=smrtmart_db

# Server
PORT=8080
GIN_MODE=release
CORS_ORIGINS=https://smrtmart.com

# Security
JWT_SECRET=your-super-secret-jwt-key

# Stripe
STRIPE_SECRET_KEY=sk_live_your_stripe_key
STRIPE_WEBHOOK_SECRET=whsec_your_webhook_secret

# File Upload
UPLOAD_PATH=./uploads
MAX_UPLOAD_SIZE=10485760  # 10MB
```

## 🚀 Deployment

### Cloudflare Deployment

1. **Build production image**:
   ```bash
   docker build -t smrtmart-api:latest .
   ```

2. **Deploy to Cloudflare Workers/Pages** (if using serverless):
   ```bash
   # Configure Cloudflare deployment
   # Set environment variables in Cloudflare dashboard
   ```

3. **Traditional server deployment**:
   ```bash
   # Deploy to your server with Docker
   docker run -d \
     --name smrtmart-api \
     -p 8080:8080 \
     --env-file .env \
     smrtmart-api:latest
   ```

### Database Migration in Production

```bash
# Run migrations on production database
migrate -path migrations \
  -database "postgres://user:pass@host:5432/dbname?sslmode=require" \
  up
```

## 🧪 Testing

```bash
# Run all tests
make test

# Run tests with coverage
go test -v -cover ./...

# Run specific test
go test -v ./internal/service/
```

## 📊 Monitoring & Health Checks

- **Health Check**: `GET /health`
- **Metrics**: Built-in request logging and error tracking
- **Database Health**: Connection pool monitoring
- **Redis Health**: Cache connectivity checks

## 🔒 Security Features

- **JWT Authentication**: Secure token-based authentication
- **Rate Limiting**: Prevent API abuse
- **CORS Protection**: Configurable cross-origin policies
- **Security Headers**: XSS, CSRF, and other security headers
- **Input Validation**: Comprehensive request validation
- **SQL Injection Prevention**: Parameterized queries

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Commit changes: `git commit -m 'Add amazing feature'`
4. Push to branch: `git push origin feature/amazing-feature`
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

For support and questions:
- Email: support@smrtmart.com
- Documentation: https://docs.smrtmart.com
- Issues: GitHub Issues

## 🎯 Roadmap

- [ ] Advanced analytics dashboard
- [ ] Multi-language support
- [ ] Advanced search with Elasticsearch
- [ ] Real-time notifications
- [ ] Mobile app API extensions
- [ ] AI-powered product recommendations
- [ ] Advanced vendor tools
- [ ] Marketing automation features