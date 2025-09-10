package main

import (
	"context"
	"fmt"
	"log"
	"github.com/jackc/pgx/v5"
	"github.com/google/uuid"
)

func main() {
	// Database connection
	dbURL := "postgresql://postgres.mqkoydypybxgcwxioqzc:admin@aws-0-eu-north-1.pooler.supabase.com:6543/postgres"
	
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	// Default vendor ID from existing data
	vendorID := "550e8400-e29b-41d4-a716-446655440002"

	// Products to add
	products := []struct {
		name         string
		description  string
		price        float64
		comparePrice float64
		category     string
		tags         []string
		images       []string
		stock        int
		featured     bool
		sku          string
	}{
		{
			name: "Braided Magnetic Charging Cable with Organizer Clip",
			description: `Premium Quality braided design ensures long-lasting performance and tangle-free use. Features a Magnetic Organizer Clip to keep your cable neatly coiled and portable. Universal Compatibility works seamlessly with USB-C and Lightning devices. Supports Fast Charging & Data Transfer with high-speed charging and reliable data transfer capabilities. 

Why Choose Us? Affordable price - cheaper than Apple.com/se, save money without compromising on quality! Convenient design perfect for travel, work, or home use. Eco-Friendly packaging thoughtfully packed to reduce waste.

Specifications: Available in 1m and 2m lengths. Made from durable braided nylon. Available connector types: USB-C and Lightning. Color options: Grey and Black.`,
			price:        249,
			comparePrice: 349,
			category:     "electronics",
			tags:         []string{"cable", "charging", "magnetic", "braided", "usb-c", "lightning"},
			images:       []string{"magnetic-charging-cable.jpg"},
			stock:        100,
			featured:     true,
			sku:          "MAG-CABLE-001",
		},
		{
			name: "Dell Thunderbolt 5/USB4 Full-Featured Dual Type-C Data Cable",
			description: `Original Dell Thunderbolt 5/USB4 Full-Featured Dual Type-C Data Cable with 240W PD Fast Charging, 8K Projection, and Emark Chip. Brand new and unopened, available in 1m, 1.5m, and 1.8m lengths with a 5.0mm diameter, making it very sturdy and durable. 

Supports multiple protocols for perfect compatibility with Thunderbolt 4, Thunderbolt 3, and USB4. Features 240W Power Delivery fast charging and 8K video output support. Professional-grade cable suitable for demanding applications.`,
			price:        199,
			comparePrice: 299,
			category:     "electronics",
			tags:         []string{"thunderbolt", "usb4", "dell", "cable", "240w", "8k", "type-c"},
			images:       []string{"dell-thunderbolt-cable.jpg"},
			stock:        50,
			featured:     true,
			sku:          "DELL-TB5-001",
		},
		{
			name: "Apple iPhone 13/13 Pro MagSafe Liquid Silicone Case",
			description: `Official Website Top-Level Animated Edition! Apple 13/13 Pro phone case with liquid silicone and magnetic technology. Compatible with iPhone 13 Promax magnetic protective case, MagSafe magnetic 15 Promax protective case, and 15 Plus models.

Features new colors in a thin and light design. Made from high-quality liquid silicone with official Apple logo. Includes charging animation pop-up window! Supports wireless charging, wallet, and power bank magnetic attachments. No color difference - the highest quality available from Huaqiangbei market.

Available for iPhone 13, 13 Pro, 13 Pro Max, 15 Plus, and 15 Pro Max models.`,
			price:        299,
			comparePrice: 499,
			category:     "smartphones",
			tags:         []string{"iphone", "case", "magsafe", "silicone", "apple", "protective"},
			images:       []string{"iphone-magsafe-case.jpg"},
			stock:        200,
			featured:     true,
			sku:          "APPLE-CASE-001",
		},
		{
			name: "Apple 29W USB-C Power Adapter A1534",
			description: `Brand New 29W Power Adapter for Apple Laptop, Type-C interface with 1-year warranty. Includes cable. Physical store clearance sale.

Compatible models:
iPad models:
- 11-inch iPad Pro (2018) with full screen
- 12.9-inch iPad Pro (2018) with full screen

Mac models:
- 12-inch MacBook
- 13-inch MacBook Air (A1534 and A1932)

This high-quality Apple power adapter ensures fast and efficient charging for your devices with a reliable 1-year warranty, providing peace of mind and convenience.`,
			price:        399,
			comparePrice: 599,
			category:     "computers",
			tags:         []string{"apple", "charger", "power-adapter", "29w", "usb-c", "macbook"},
			images:       []string{"apple-29w-adapter.jpg"},
			stock:        75,
			featured:     false,
			sku:          "APPLE-29W-001",
		},
		{
			name: "MacBook Air M3 13-inch Protective Case - Grass Green",
			description: `Compatible with the 2024 model M3 Air 13-inch laptop protective case. The Grass Green White color case fits perfectly with your laptop, providing a comfortable grip and a woven outer shell. 

This case not only offers excellent protection but also adds a stylish touch to your device. Precision cutouts ensure easy access to all ports and features. Lightweight design doesn't add bulk to your MacBook Air.`,
			price:        199,
			comparePrice: 299,
			category:     "computers",
			tags:         []string{"macbook", "case", "m3", "protective", "laptop", "accessories"},
			images:       []string{"macbook-air-case-green.jpg"},
			stock:        120,
			featured:     false,
			sku:          "MAC-CASE-M3-001",
		},
		{
			name: "MacBook Pro MagSafe 3 Charging Cable - Midnight Blue",
			description: `MacBook Pro MagSafe Charging Cable for M4, M2, and M3 chips. This magnetic charging cable is designed for MacBook Air 14 and 16-inch models and is an original USB-C to MagSafe 3 connecting line. 

Compatible with MacBook Air notebooks and Apple Air A2941 computer chargers (A2991, A2992, A2918, A3114, A3113). Features the convenient MagSafe magnetic connection that prevents accidental disconnection. Available in elegant Midnight Blue color.`,
			price:        599,
			comparePrice: 799,
			category:     "computers",
			tags:         []string{"magsafe", "macbook", "charging", "cable", "m3", "m2", "m4"},
			images:       []string{"magsafe3-cable-blue.jpg"},
			stock:        80,
			featured:     true,
			sku:          "MAGSAFE3-001",
		},
		{
			name: "Huawei GT2 Pro Smart Watch - Phantom Black",
			description: `Huawei GT2 Pro Smart Watch in Phantom Black - a premium health monitoring smartwatch. Brand new, opened only for inspection. Fully functional with all accessories included.

Features comprehensive health monitoring including heart rate, SpO2, sleep tracking, and stress monitoring. Premium titanium body with sapphire glass display. Supports wireless charging and offers up to 14 days battery life. Water resistant up to 5ATM. Includes GPS, GLONASS, and multiple sport modes.`,
			price:        1999,
			comparePrice: 2499,
			category:     "wearables",
			tags:         []string{"smartwatch", "huawei", "gt2-pro", "health", "fitness", "wearable"},
			images:       []string{"huawei-gt2-pro.jpg"},
			stock:        30,
			featured:     true,
			sku:          "HUAWEI-GT2P-001",
		},
		{
			name: "Mtag Apple AirTag Compatible Tracker",
			description: `Mtag Apple AirTag compatible alternative from Diwenjia! Perfect for locating lost bicycles and tracking personal belongings. This device is a must-have for ensuring your valuables stay safe and sound. 

With its innovative features and full compatibility with Apple's Find My network, you'll never lose track of your belongings again. Long battery life, water-resistant design, and precision finding capabilities. Works seamlessly with iPhone, iPad, and Mac devices.`,
			price:        149,
			comparePrice: 249,
			category:     "electronics",
			tags:         []string{"tracker", "airtag", "compatible", "finder", "location", "bluetooth"},
			images:       []string{"mtag-tracker.jpg"},
			stock:        150,
			featured:     false,
			sku:          "MTAG-001",
		},
	}

	// Insert products
	for _, p := range products {
		query := `
			INSERT INTO products (
				id, vendor_id, name, description, price, compare_price, 
				category, tags, images, stock, status, featured, sku
			) VALUES (
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
			)`

		productID := uuid.New()
		
		_, err = conn.Exec(context.Background(), query,
			productID,
			vendorID,
			p.name,
			p.description,
			p.price,
			p.comparePrice,
			p.category,
			p.tags,
			p.images,
			p.stock,
			"active",
			p.featured,
			p.sku,
		)
		
		if err != nil {
			log.Printf("Error inserting product %s: %v\n", p.name, err)
		} else {
			fmt.Printf("Successfully added product: %s (ID: %s)\n", p.name, productID)
		}
	}

	fmt.Println("\nAll products have been processed!")
}