package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/mauricio3g/golang-fiber/database"
	"net/http"
	"os"
	"github.com/Chris-cez/BaseShopSystem/storage"
	"github.com/Chris-cez/BaseShopSystem/models"
)

type Product struct {
	Code  string		`json:"code"`
	Price float32 		`json:"price"`
	Name  string		`json:"name"`
	NCM   string		`json:"ncm"`
	UM    string		`json:"um"`
	Description string	`json:"description"`
}

func (r *Repository) CreateProduct(c *fiber.Ctx) error {
	product := Product{}

	err := context.BodyParser(&product)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
		return err
	}
	err := r.DB.Create(&product).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create product"})
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"product": "product has been added"})
	return nil
}

func (r *Repository) GetProducts(c *fiber.Ctx) error {
	productModels := []models.Product{}
	err := r.DB.Find(&productModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get products"})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "products fetched successfully", 
		"data": productModels})
	return nil
}

type Repository struct {
	DB *gorm.DB
}

func(r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/products", r.CreateProduct)
	api.Get("/products", r.GetProducts)
	api.Get("/products/:id", r.GetProductByID)
	api.Put("/products/:id", r.UpdateProduct)
	api.Delete("/products/:id", r.DeleteProduct)
}

func main()  {
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
	r := Repository{DB: db}

	app := fiber.New()
	database.ConnectDB()
	routes.SetupRoutes(app)
	app.Listen(":8080")
}


