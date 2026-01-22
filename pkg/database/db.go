package database

import (
	"log"

	"github.com/ahmadeko2017/backed-golang-tugas-1/internal/entity"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Use SQLite for simplified local development
	dsn := "tugas1.db"

	var err error
	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Auto Migrate
	err = DB.AutoMigrate(&entity.Category{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	log.Println("Database connected and migrated successfully (SQLite)")
}
