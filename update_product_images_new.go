package main

import (
	"context"
	"fmt"
	"log"
	"github.com/jackc/pgx/v5"
)

func main() {
	// Database connection
	dbURL := "postgresql://postgres.mqkoydypybxgcwxioqzc:admin@aws-0-eu-north-1.pooler.supabase.com:6543/postgres"
	
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	// Map product names to their actual image files
	productImageMappings := []struct {
		productName string
		imageFile   string
	}{
		{
			productName: "Braided Magnetic Charging Cable with Organizer Clip",
			imageFile:   "usb c iphone cable.jpg",
		},
		{
			productName: "Dell Thunderbolt 5/USB4 Full-Featured Dual Type-C Data Cable",
			imageFile:   "8k data cable dell.jpg",
		},
		{
			productName: "Apple iPhone 13/13 Pro MagSafe Liquid Silicone Case",
			imageFile:   "iphone16 promaxcase.jpg",
		},
		{
			productName: "Apple 29W USB-C Power Adapter A1534",
			imageFile:   "macbookair adaptor and cable.png",
		},
		{
			productName: "MacBook Air M3 13-inch Protective Case - Grass Green",
			imageFile:   "macbookair m3 weaving case.jpg",
		},
		{
			productName: "MacBook Pro MagSafe 3 Charging Cable - Midnight Blue",
			imageFile:   "macbook m4 charging cable.png",
		},
		{
			productName: "Huawei GT2 Pro Smart Watch - Phantom Black",
			imageFile:   "huaweismartwatch.jpg",
		},
		{
			productName: "Mtag Apple AirTag Compatible Tracker",
			imageFile:   "mtrackingtag.jpg",
		},
	}

	// Update each product's image
	for _, mapping := range productImageMappings {
		query := `UPDATE products SET images = ARRAY[$1]::text[] WHERE name = $2`
		
		result, err := conn.Exec(context.Background(), query, mapping.imageFile, mapping.productName)
		if err != nil {
			log.Printf("Error updating image for %s: %v\n", mapping.productName, err)
			continue
		}
		
		rowsAffected := result.RowsAffected()
		if rowsAffected > 0 {
			fmt.Printf("✓ Updated image for: %s -> %s\n", mapping.productName, mapping.imageFile)
		} else {
			fmt.Printf("✗ No product found with name: %s\n", mapping.productName)
		}
	}

	fmt.Println("\nAll product images have been updated!")
}