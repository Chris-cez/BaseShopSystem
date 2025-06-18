package main

import (
	"log"
	"os"

	"github.com/Chris-cez/BaseShopSystem/models"
	"github.com/Chris-cez/BaseShopSystem/routes"
	"github.com/Chris-cez/BaseShopSystem/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func insertTestData(db *gorm.DB) error {
	// Verifica se já existe o endereço de teste
	var count int64
	db.Model(&models.Address{}).Where("cep = ?", "12345678").Count(&count)
	if count == 0 {
		address := models.Address{
			Logradouro:  "Rua Teste",
			Numero:      "123",
			Complemento: "Apto 1",
			Bairro:      "Centro",
			Municipio:   "Cidade Teste",
			UF:          "SP",
			CEP:         "12345678",
		}
		if err := db.Create(&address).Error; err != nil {
			return err
		}

		// Cria empresa de teste associada ao endereço
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("senha123"), bcrypt.DefaultCost)
		company := models.Company{
			Name:              "Empresa Teste",
			CNPJ:              "12345678000199",
			InscricaoEstadual: "ISENTO",
			Password:          string(hashedPassword),
			Address_id:        int(address.ID),
		}
		if err := db.Create(&company).Error; err != nil {
			return err
		}
	}
	return nil
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
