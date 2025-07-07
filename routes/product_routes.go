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

// CreateProduct godoc
// @Summary Cria um novo produto
// @Description Adiciona um novo produto ao banco de dados
// @Tags product
// @Accept  json
// @Produce  json
// @Param product body models.Product true "Dados do produto"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /api/products [post]
func (r *ProductRepository) CreateProduct(c *fiber.Ctx) error {
	product := models.Product{}

	err := c.BodyParser(&product)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
		return err
	}
	if product.Price <= 0 {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "product price must be greater than zero"})
		return nil
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

// GetProducts godoc
// @Summary Lista todos os produtos
// @Description Retorna todos os produtos cadastrados
// @Tags product
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/products [get]
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

// GetProductByID godoc
// @Summary Busca produto por ID
// @Description Retorna um produto pelo seu ID
// @Tags product
// @Accept  json
// @Produce  json
// @Param id path int true "ID do produto"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/products/{id} [get]
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

// GetProductsByName godoc
// @Summary Busca produtos por nome
// @Description Retorna produtos que contenham o nome informado
// @Tags product
// @Accept  json
// @Produce  json
// @Param name path string true "Nome do produto"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/products/name/{name} [get]
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

// GetProductsByClass godoc
// @Summary Busca produtos por classe
// @Description Retorna produtos de uma determinada classe
// @Tags product
// @Accept  json
// @Produce  json
// @Param class_id path int true "ID da classe"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/products/class/{class_id} [get]
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

// UpdateProduct godoc
// @Summary Atualiza um produto
// @Description Atualiza os dados de um produto existente
// @Tags product
// @Accept  json
// @Produce  json
// @Param id path int true "ID do produto"
// @Param product body models.Product true "Dados do produto"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/products/{id} [put]
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

	if updateData.Price <= 0 {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "product price must be greater than zero"})
		return nil
	}
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

// DeleteProduct godoc
// @Summary Remove um produto
// @Description Deleta um produto pelo ID
// @Tags product
// @Accept  json
// @Produce  json
// @Param id path int true "ID do produto"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/products/{id} [delete]
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
	api := app.Group("/api", middleware.AuthRequired)
	api.Post("/products", r.CreateProduct)
	api.Get("/products", r.GetProducts)
	api.Get("/products/:id", r.GetProductByID)
	api.Get("/products/name/:name", r.GetProductsByName)
	api.Get("/products/class/:class_id", r.GetProductsByClass)
	api.Put("/products/:id", r.UpdateProduct)
	api.Delete("/products/:id", r.DeleteProduct)
}
