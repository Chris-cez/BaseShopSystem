package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Chris-cez/BaseShopSystem/models"
	"github.com/Chris-cez/BaseShopSystem/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Product struct {
	Code        string  `json:"code"`
	Price       float32 `json:"price"`
	Name        string  `json:"name"`
	NCM         string  `json:"ncm"`
	UM          string  `json:"um"`
	Description string  `json:"description"`
}

func (r *Repository) CreateProduct(c *fiber.Ctx) error {
	product := Product{}

	err := c.BodyParser(&product)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
		return err
	}
	err = r.DB.Create(&product).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create product"})
		return err
	}

	c.Status(http.StatusOK).JSON(
		&fiber.Map{"product": "product has been added"})
	return nil
}

func (r *Repository) GetProducts(c *fiber.Ctx) error {
	productModels := []models.Product{}
	err := r.DB.Find(&productModels).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get products"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "products fetched successfully",
			"data": productModels})
	return nil
}

func (r *Repository) DeleteProduct(c *fiber.Ctx) error {
	productModel := models.Product{}

	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id can not be empty"})
		return nil
	}
	err := r.DB.Delete(productModel, id).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not delete product"})
		return err
	}

	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "product deleted successfully"})
	return nil
}

func (r *Repository) GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	productModel := models.Product{}

	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id can not be empty"})
		return nil
	}
	fmt.Println("The ID is", id)

	err := r.DB.Where("id = ?", id).First(productModel).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the product"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "product fetched successfully",
			"data": productModel})
	return nil
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/products", r.CreateProduct)
	api.Get("/products", r.GetProducts)
	api.Get("/products/:id", r.GetProductByID)
	api.Delete("/products/:id", r.DeleteProduct)
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

	err = models.MigrateProducts(db)
	if err != nil {
		log.Fatal("Error migrating products\n", err)
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	log.Fatal(app.Listen(":8080"))
}
