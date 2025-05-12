package database

import (
	"game_log_hub/models"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Initialize initializes the database connection
func Initialize() {
	var err error

	// Connect to SQLite database
	DB, err = gorm.Open(sqlite.Open("game_log_hub.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Database connection established")

	// Auto migrate the database schema
	err = DB.AutoMigrate(&models.LoginError{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	log.Println("Database migration completed")
}
