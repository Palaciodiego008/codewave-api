package config

import (
	"codewave/models"
	"fmt"
	"log"
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

	// Conectar a la base de datos
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "public.",
		},
	})
	if err != nil {
		panic("Failed to connect to database!")
	}

	// Migrar los modelos si las tablas no existen
	if !DB.Migrator().HasTable(&models.User{}) || !DB.Migrator().HasTable(&models.Project{}) {
		modelsToMigrate := []interface{}{&models.User{}, &models.Project{}}
		err = DB.AutoMigrate(modelsToMigrate...)
		if err != nil {
			panic("Failed to migrate the schema!")
		}
	} else {
		log.Println("Database tables already exist. Skipping migration.")
	}
}
