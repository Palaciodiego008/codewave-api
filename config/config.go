package config

import (
	"codewave/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() {

	var err error
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	fmt.Println("Error loading .env file")
	// }

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "postgres"
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "123456"
	}
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}
	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "codewave"
	}

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		user, password, host, port, dbname)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "public.",
		},
	})
	if err != nil {
		panic("Failed to connect to database!")
	}

	modelsToMigrate := []interface{}{&models.User{}, &models.Project{}}
	err = DB.AutoMigrate(modelsToMigrate...)
	if err != nil {
		panic("Failed to migrate the schema!")
	}

}
