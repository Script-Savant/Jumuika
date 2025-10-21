package config

import (
	"Jumuika/common/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnection() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load env variables")
	}

	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to make a database connection")
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Category{},
		&models.Location{},
		&models.Meeting{},
		&models.RSVP{},
		&models.Comment{},
		&models.Friendship{},
	); err != nil {
		log.Fatal("Failed to migrate tables to the database")
	}

	log.Println("Database connection established successfully")

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
