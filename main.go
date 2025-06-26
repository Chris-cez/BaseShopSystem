package main

import (
	"log"
	"os"

	"github.com/Chris-cez/BaseShopSystem/routes"
	"github.com/Chris-cez/BaseShopSystem/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file\n", err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("Error connecting to database\n", err)
	}

	err = storage.MigrateModels(db)
	if err != nil {
		log.Fatal("Error migrating models\n", err)
	}

	// Inserir dados de teste
	if err := insertTestData(db); err != nil {
		log.Fatal("Error inserting test data\n", err)
	}

	app := fiber.New()

	routes.SetupRoutes(app, db)

	log.Fatal(app.Listen(":8080"))
}
