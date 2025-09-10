package database

import (
	"database/sql"
	"embed"
	"fmt"
	"net/url"
	"os"
	"strings"

	"smrtmart-go-postgresql/internal/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)


func Initialize(cfg config.DatabaseConfig) (*sql.DB, error) {
	var dsn string
	
	// Check if JAWSDB_URL is available (Heroku MySQL addon)
	if jawsDBURL := os.Getenv("JAWSDB_URL"); jawsDBURL != "" {
		// JawsDB URL format: mysql://username:password@hostname:port/database
		// Convert to MySQL DSN format: username:password@tcp(hostname:port)/database
		dsn = convertJawsDBURLToMySQLDSN(jawsDBURL)
	} else {
		// MySQL DSN format: user:password@tcp(host:port)/database?parseTime=true
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	return db, nil
}

func RunMigrations(cfg config.DatabaseConfig, migrationFS embed.FS) error {
	var dsn string
	
	// Check if JAWSDB_URL is available (Heroku MySQL addon)
	if jawsDBURL := os.Getenv("JAWSDB_URL"); jawsDBURL != "" {
		// JawsDB URL format: mysql://username:password@hostname:port/database
		// Convert to MySQL DSN format: username:password@tcp(hostname:port)/database
		dsn = convertJawsDBURLToMySQLDSN(jawsDBURL)
	} else {
		// MySQL DSN format: user:password@tcp(host:port)/database?parseTime=true
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database for migrations: %w", err)
	}
	defer db.Close()

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("failed to create mysql driver: %w", err)
	}

	// Use embedded migration files
	sourceDriver, err := iofs.New(migrationFS, "migrations")
	if err != nil {
		return fmt.Errorf("failed to create embedded migration source: %w", err)
	}
	
	m, err := migrate.NewWithSourceInstance("iofs", sourceDriver, "mysql", driver)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

// convertJawsDBURLToMySQLDSN converts JawsDB URL to MySQL DSN format
// Input: mysql://username:password@hostname:port/database
// Output: username:password@tcp(hostname:port)/database?parseTime=true
func convertJawsDBURLToMySQLDSN(jawsDBURL string) string {
	parsedURL, err := url.Parse(jawsDBURL)
	if err != nil {
		// Return empty string if parsing fails
		return ""
	}

	password, _ := parsedURL.User.Password()
	username := parsedURL.User.Username()
	hostname := parsedURL.Hostname()
	port := parsedURL.Port()
	database := strings.TrimPrefix(parsedURL.Path, "/")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		username, password, hostname, port, database)
}