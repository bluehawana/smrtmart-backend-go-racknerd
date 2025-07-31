package main

import (
	"fmt"
	"os"
	"smrtmart-go-postgresql/internal/config"
)

func main() {
	cfg := config.Load()
	fmt.Printf("DB Config: Host=%s, Port=%s, User=%s, Password=%s, Name=%s, SSLMode=%s\n",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name, cfg.Database.SSLMode)
	
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name, cfg.Database.SSLMode)
	
	fmt.Printf("DSN: %s\n", dsn)
	
	// Print all environment variables starting with DB_
	for _, env := range os.Environ() {
		if len(env) > 3 && env[:3] == "DB_" {
			fmt.Printf("ENV: %s\n", env)
		}
	}
}