package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	// Database connection string
	dbURL := "postgresql://postgres:admin@db.mqkoydypybxgcwxioqzc.supabase.co:5432/postgres"
	
	// Connect to database
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	// Update Dell XPS 15 Developer Edition image
	query := `UPDATE products 
			  SET images = ARRAY['https://mqkoydypybxgcwxioqzc.supabase.co/storage/v1/object/public/products/dell-xps-15-2023.jpg']
			  WHERE name = 'Dell XPS 15 Developer Edition'`

	result, err := conn.Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Unable to update product: %v\n", err)
	}

	fmt.Printf("Successfully updated %d rows\n", result.RowsAffected())

	// Verify the update
	var name string
	var images []string
	err = conn.QueryRow(context.Background(), 
		"SELECT name, images FROM products WHERE name = 'Dell XPS 15 Developer Edition'").Scan(&name, &images)
	if err != nil {
		log.Fatalf("Unable to query product: %v\n", err)
	}

	fmt.Printf("Product: %s\n", name)
	fmt.Printf("Images: %v\n", images)
}