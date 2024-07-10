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
	// Cargar variables del archivo .env
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	// Obtener las variables de entorno
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	// Construir la cadena de conexión
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=require",
		user, password, host, port, dbname)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "public.",
		},
	})
	if err != nil {
		panic("Failed to connect to database!")
	}

	// Migrar los modelos
	modelsToMigrate := []interface{}{&models.User{}, &models.Project{}}
	err = DB.AutoMigrate(modelsToMigrate...)
	if err != nil {
		panic("Failed to migrate the schema!")
	}
}
