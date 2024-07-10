package config

import (
	"codewave/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() {
	// Load the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dsn := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "public.",
		},
	})
	if err != nil {
		panic("Failed to connect to database!")
	}
	// Migrate the schema
	modelsToMigrate := []interface{}{&models.User{}, &models.Project{}}
	err = DB.AutoMigrate(modelsToMigrate...)
	if err != nil {
		panic("Failed to migrate the schema!")
	}
}
