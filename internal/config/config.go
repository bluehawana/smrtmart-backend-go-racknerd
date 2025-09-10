package config

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	JWT      JWTConfig
	Stripe   StripeConfig
	Upload   UploadConfig
	Email    EmailConfig
	Redis    RedisConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type ServerConfig struct {
	Port        string
	Mode        string
	CORSOrigins []string
}

type JWTConfig struct {
	Secret string
}

type StripeConfig struct {
	SecretKey     string
	WebhookSecret string
}

type UploadConfig struct {
	Path        string
	MaxFileSize int64
}

type EmailConfig struct {
	SMTPHost string
	SMTPPort string
	User     string
	Password string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

func Load() *Config {
	var dbConfig DatabaseConfig
	
	// Check for DATABASE_URL first (Heroku PostgreSQL)
	if databaseURL := getEnv("DATABASE_URL", ""); databaseURL != "" {
		var err error
		dbConfig, err = parsePostgreSQLURL(databaseURL)
		if err != nil {
			// Fallback to individual env vars if URL parsing fails
			dbConfig = DatabaseConfig{
				Host:     getEnv("DB_HOST", "localhost"),
				Port:     getEnv("DB_PORT", "5432"),
				User:     getEnv("DB_USER", "postgres"),
				Password: getEnv("DB_PASSWORD", ""),
				Name:     getEnv("DB_NAME", "postgres"),
				SSLMode:  getEnv("DB_SSLMODE", "require"),
			}
		}
	} else {
		// Use individual environment variables
		dbConfig = DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "postgres"),
			SSLMode:  getEnv("DB_SSLMODE", "require"),
		}
	}
	
	return &Config{
		Database: dbConfig,
		Server: ServerConfig{
			Port:        getEnv("PORT", "8080"),
			Mode:        getEnv("GIN_MODE", "debug"),
			CORSOrigins: strings.Split(getEnv("CORS_ORIGINS", "http://localhost:3000"), ","),
		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", "your-secret-key"),
		},
		Stripe: StripeConfig{
			SecretKey:     getEnv("STRIPE_SECRET_KEY", ""),
			WebhookSecret: getEnv("STRIPE_WEBHOOK_SECRET", ""),
		},
		Upload: UploadConfig{
			Path:        getEnv("UPLOAD_PATH", "./uploads"),
			MaxFileSize: getEnvAsInt64("MAX_UPLOAD_SIZE", 10485760), // 10MB
		},
		Email: EmailConfig{
			SMTPHost: getEnv("SMTP_HOST", "smtp.gmail.com"),
			SMTPPort: getEnv("SMTP_PORT", "587"),
			User:     getEnv("SMTP_USER", ""),
			Password: getEnv("SMTP_PASS", ""),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt64(key string, defaultValue int64) int64 {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// parsePostgreSQLURL parses a PostgreSQL URL into DatabaseConfig
// Format: postgresql://username:password@hostname:port/database
func parsePostgreSQLURL(databaseURL string) (DatabaseConfig, error) {
	parsedURL, err := url.Parse(databaseURL)
	if err != nil {
		return DatabaseConfig{}, fmt.Errorf("failed to parse DATABASE_URL: %w", err)
	}

	password, _ := parsedURL.User.Password()
	port := parsedURL.Port()
	if port == "" {
		port = "5432" // Default PostgreSQL port
	}

	// Extract SSL mode from query parameters
	sslMode := "require" // Default for hosted PostgreSQL
	if sslModeParam := parsedURL.Query().Get("sslmode"); sslModeParam != "" {
		sslMode = sslModeParam
	}
	
	return DatabaseConfig{
		Host:     parsedURL.Hostname(),
		Port:     port,
		User:     parsedURL.User.Username(),
		Password: password,
		Name:     strings.TrimPrefix(parsedURL.Path, "/"),
		SSLMode:  sslMode,
	}, nil
}