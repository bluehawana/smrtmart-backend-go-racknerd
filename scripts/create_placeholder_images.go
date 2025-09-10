package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func main() {
	// Product images to create
	products := map[string]string{
		"magnetic-charging-cable.jpg": "Magnetic\nCharging\nCable",
		"dell-thunderbolt-cable.jpg":  "Dell\nThunderbolt\nCable",
		"iphone-magsafe-case.jpg":     "iPhone\nMagSafe\nCase",
		"apple-29w-adapter.jpg":       "Apple 29W\nAdapter",
		"macbook-air-case-green.jpg":  "MacBook\nCase\nGreen",
		"magsafe3-cable-blue.jpg":     "MagSafe 3\nCable",
		"huawei-gt2-pro.jpg":          "Huawei\nGT2 Pro\nWatch",
		"mtag-tracker.jpg":            "Mtag\nTracker",
	}

	// Create uploads directory if it doesn't exist
	uploadsDir := "./uploads"
	if err := os.MkdirAll(uploadsDir, 0755); err != nil {
		fmt.Printf("Error creating uploads directory: %v\n", err)
		return
	}

	// Colors for different product types
	colors := []color.RGBA{
		{R: 51, G: 51, B: 51, A: 255},    // Dark gray
		{R: 0, G: 122, B: 255, A: 255},   // Blue
		{R: 52, G: 199, B: 89, A: 255},   // Green
		{R: 255, G: 59, B: 48, A: 255},   // Red
		{R: 90, G: 200, B: 250, A: 255},  // Light blue
		{R: 255, G: 149, B: 0, A: 255},   // Orange
		{R: 88, G: 86, B: 214, A: 255},   // Purple
		{R: 50, G: 173, B: 230, A: 255},  // Sky blue
	}

	i := 0
	for filename, text := range products {
		// Create 800x800 image
		img := image.NewRGBA(image.Rect(0, 0, 800, 800))
		
		// Fill with color
		bgColor := colors[i%len(colors)]
		draw.Draw(img, img.Bounds(), &image.Uniform{bgColor}, image.Point{}, draw.Src)

		// Add white rectangle in center
		centerRect := image.Rect(100, 250, 700, 550)
		white := color.RGBA{R: 255, G: 255, B: 255, A: 255}
		draw.Draw(img, centerRect, &image.Uniform{white}, image.Point{}, draw.Src)

		// Add text
		addLabel(img, text, bgColor)

		// Save image
		filepath := filepath.Join(uploadsDir, filename)
		file, err := os.Create(filepath)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", filename, err)
			continue
		}
		defer file.Close()

		err = jpeg.Encode(file, img, &jpeg.Options{Quality: 80})
		if err != nil {
			fmt.Printf("Error encoding image %s: %v\n", filename, err)
			continue
		}

		fmt.Printf("Created placeholder image: %s\n", filepath)
		i++
	}

	fmt.Println("\nAll placeholder images created successfully!")
}

func addLabel(img *image.RGBA, label string, textColor color.RGBA) {
	point := fixed.Point26_6{X: fixed.Int26_6(400 * 64), Y: fixed.Int26_6(400 * 64)}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(textColor),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	
	// Split text by newlines and center each line
	lines := strings.Split(label, "\n")
	lineHeight := 20
	totalHeight := len(lines) * lineHeight
	startY := 400 - totalHeight/2
	
	for i, line := range lines {
		// Calculate text width for centering
		textWidth := len(line) * 7 // Approximate width per character
		x := 400 - textWidth/2
		y := startY + i*lineHeight
		
		d.Dot = fixed.Point26_6{
			X: fixed.Int26_6(x * 64),
			Y: fixed.Int26_6(y * 64),
		}
		d.DrawString(line)
	}
}