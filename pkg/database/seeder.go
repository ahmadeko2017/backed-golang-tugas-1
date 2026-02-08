package database

import (
	"log"

	"github.com/ahmadeko2017/backed-golang-tugas/internal/entity"
)

// SeedData inserts sample categories and products into the database
func SeedData() {
	// Check if data already exists
	var categoryCount int64
	DB.Model(&entity.Category{}).Count(&categoryCount)

	if categoryCount > 0 {
		log.Println("Database already has data, skipping seeding")
		return
	}

	// Sample Categories
	categories := []entity.Category{
		{Name: "Electronics", Description: "Electronic devices and accessories"},
		{Name: "Clothing", Description: "Apparel and fashion items"},
		{Name: "Books", Description: "Physical and digital books"},
		{Name: "Home & Garden", Description: "Home improvement and gardening supplies"},
		{Name: "Sports", Description: "Sports equipment and gear"},
	}

	for i := range categories {
		if err := DB.Create(&categories[i]).Error; err != nil {
			log.Printf("Failed to seed category %s: %v", categories[i].Name, err)
		}
	}

	log.Println("Categories seeded successfully")

	// Sample Products
	products := []entity.Product{
		// Electronics
		{Name: "Laptop Dell XPS 13", Description: "High-performance ultrabook", Price: 1299.99, Stock: 15, CategoryID: categories[0].ID},
		{Name: "Wireless Mouse Logitech", Description: "Ergonomic wireless mouse", Price: 29.99, Stock: 50, CategoryID: categories[0].ID},
		{Name: "USB-C Hub", Description: "7-in-1 USB-C adapter", Price: 49.99, Stock: 30, CategoryID: categories[0].ID},

		// Clothing
		{Name: "Cotton T-Shirt", Description: "Premium quality cotton t-shirt", Price: 19.99, Stock: 100, CategoryID: categories[1].ID},
		{Name: "Denim Jeans", Description: "Classic blue denim jeans", Price: 59.99, Stock: 45, CategoryID: categories[1].ID},
		{Name: "Leather Jacket", Description: "Genuine leather jacket", Price: 199.99, Stock: 20, CategoryID: categories[1].ID},

		// Books
		{Name: "Clean Code", Description: "A Handbook of Agile Software Craftsmanship", Price: 42.99, Stock: 25, CategoryID: categories[2].ID},
		{Name: "The Pragmatic Programmer", Description: "Your Journey to Mastery", Price: 44.99, Stock: 30, CategoryID: categories[2].ID},
		{Name: "Design Patterns", Description: "Elements of Reusable Object-Oriented Software", Price: 54.99, Stock: 18, CategoryID: categories[2].ID},

		// Home & Garden
		{Name: "LED Desk Lamp", Description: "Adjustable brightness desk lamp", Price: 34.99, Stock: 40, CategoryID: categories[3].ID},
		{Name: "Garden Tool Set", Description: "10-piece gardening tool kit", Price: 79.99, Stock: 22, CategoryID: categories[3].ID},

		// Sports
		{Name: "Yoga Mat", Description: "Non-slip exercise yoga mat", Price: 24.99, Stock: 60, CategoryID: categories[4].ID},
		{Name: "Running Shoes Nike", Description: "Lightweight running shoes", Price: 89.99, Stock: 35, CategoryID: categories[4].ID},
		{Name: "Resistance Bands Set", Description: "5-piece resistance band set", Price: 19.99, Stock: 55, CategoryID: categories[4].ID},
	}

	for i := range products {
		if err := DB.Create(&products[i]).Error; err != nil {
			log.Printf("Failed to seed product %s: %v", products[i].Name, err)
		}
	}

	log.Println("Products seeded successfully")
	log.Printf("Seeded %d categories and %d products", len(categories), len(products))
}
