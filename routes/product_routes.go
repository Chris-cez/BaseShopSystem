package routes

import (
	"fmt"
	"net/http"

	"github.com/Chris-cez/BaseShopSystem/middleware"
	"github.com/Chris-cez/BaseShopSystem/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func (r *ProductRepository) CreateProduct(c *fiber.Ctx) error {
	product := models.Product{}

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

func (r *ProductRepository) GetProducts(c *fiber.Ctx) error {
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

func (r *ProductRepository) GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	productModel := models.Product{}

	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id can not be empty"})
		return nil
	}
	fmt.Println("The ID is", id)

	err := r.DB.Where("id = ?", id).First(&productModel).Error
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

func (r *ProductRepository) GetProductsByName(c *fiber.Ctx) error {
	name := c.Params("name")
	productModels := []models.Product{}

	err := r.DB.Where("name LIKE ?", "%"+name+"%").Find(&productModels).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get products by name"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "products fetched successfully",
			"data": productModels})
	return nil
}

func (r *ProductRepository) GetProductsByClass(c *fiber.Ctx) error {
	classID := c.Params("class_id")
	productModels := []models.Product{}

	err := r.DB.Where("class_id = ?", classID).Find(&productModels).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get products by class"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "products fetched successfully",
			"data": productModels})
	return nil
}

func (r *ProductRepository) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product

	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id can not be empty"})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(&product).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "product not found"})
		return err
	}

	var updateData models.Product
	err = c.BodyParser(&updateData)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
		return err
	}

	// Atualize apenas os campos relevantes
	product.Code = updateData.Code
	product.Price = updateData.Price
	product.Name = updateData.Name
	product.NCM = updateData.NCM
	product.GTIN = updateData.GTIN
	product.UM = updateData.UM
	product.Description = updateData.Description
	product.ClassID = updateData.ClassID
	product.Stock = updateData.Stock
	product.ValTrib = updateData.ValTrib

	err = r.DB.Save(&product).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not update product"})
		return err
	}

	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "product updated successfully"})
	return nil
}

func (r *ProductRepository) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product := models.Product{}

	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id can not be empty"})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(&product).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "product not found"})
		return err
	}

	err = r.DB.Delete(&product).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not delete product"})
		return err
	}

	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "product deleted successfully"})
	return nil
}

func (r *ProductRepository) SetupProductRoutes(app *fiber.App) {
	api := app.Group("/api", middleware.AuthRequired) // Adiciona o middleware aqui
	api.Post("/products", r.CreateProduct)
	api.Get("/products", r.GetProducts)
	api.Get("/products/:id", r.GetProductByID)
	api.Get("/products/name/:name", r.GetProductsByName)
	api.Get("/products/class/:class_id", r.GetProductsByClass)
	api.Put("/products/:id", r.UpdateProduct)
	api.Delete("/products/:id", r.DeleteProduct)
}
