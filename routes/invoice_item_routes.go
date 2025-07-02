package routes

import (
	"net/http"

	"github.com/Chris-cez/BaseShopSystem/middleware"
	"github.com/Chris-cez/BaseShopSystem/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type InvoiceItemRepository struct {
	DB *gorm.DB
}

// CreateInvoiceItem godoc
// @Summary Cria um novo item de nota fiscal
// @Description Adiciona um novo item à nota fiscal
// @Tags invoice_item
// @Accept  json
// @Produce  json
// @Param invoice_item body models.InvoiceItem true "Dados do item da nota"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /api/invoice_items [post]
func (r *InvoiceItemRepository) CreateInvoiceItem(c *fiber.Ctx) error {
	invoiceItem := models.InvoiceItem{}

	err := c.BodyParser(&invoiceItem)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
		return err
	}

	// Buscar preço do produto
	var product models.Product
	if err := r.DB.First(&product, invoiceItem.ProductID).Error; err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "product not found"})
		return err
	}

	invoiceItem.Price = product.Price
	invoiceItem.ValorTotal = product.Price * float64(invoiceItem.Quantity)

	// Adiciona o item
	err = r.DB.Create(&invoiceItem).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create invoice item"})
		return err
	}

	// Atualiza o total da nota fiscal
	var invoice models.Invoice
	if err := r.DB.Where("numero = ?", invoiceItem.InvoiceID).First(&invoice).Error; err == nil {
		invoice.TotalValue += invoiceItem.ValorTotal
		r.DB.Save(&invoice)
	}

	c.Status(http.StatusOK).JSON(
		&fiber.Map{"invoice_item": "invoice item has been added"})
	return nil
}

// GetInvoiceItems godoc
// @Summary Lista todos os itens de nota fiscal
// @Description Retorna todos os itens de nota fiscal cadastrados
// @Tags invoice_item
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/invoice_items [get]
func (r *InvoiceItemRepository) GetInvoiceItems(c *fiber.Ctx) error {
	invoiceItemModels := []models.InvoiceItem{}
	err := r.DB.Find(&invoiceItemModels).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get invoice items"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "invoice items fetched successfully",
			"data": invoiceItemModels})
	return nil
}

// GetInvoiceItemByID godoc
// @Summary Busca item de nota fiscal por ID
// @Description Retorna um item de nota fiscal pelo seu ID
// @Tags invoice_item
// @Accept  json
// @Produce  json
// @Param id path int true "ID do item da nota"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/invoice_items/{id} [get]
func (r *InvoiceItemRepository) GetInvoiceItemByID(c *fiber.Ctx) error {
	id := c.Params("id")
	invoiceItemModel := models.InvoiceItem{}

	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id can not be empty"})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(&invoiceItemModel).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the invoice item"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "invoice item fetched successfully",
			"data": invoiceItemModel})
	return nil
}

// UpdateInvoiceItem godoc
// @Summary Atualiza um item de nota fiscal
// @Description Atualiza os dados de um item de nota fiscal existente
// @Tags invoice_item
// @Accept  json
// @Produce  json
// @Param id path int true "ID do item da nota"
// @Param invoice_item body models.InvoiceItem true "Dados do item da nota"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/invoice_items/{id} [put]
func (r *InvoiceItemRepository) UpdateInvoiceItem(c *fiber.Ctx) error {
	id := c.Params("id")
	invoiceItem := models.InvoiceItem{}

	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id can not be empty"})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(&invoiceItem).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "invoice item not found"})
		return err
	}

	err = c.BodyParser(&invoiceItem)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
		return err
	}

	err = r.DB.Save(&invoiceItem).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not update invoice item"})
		return err
	}

	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "invoice item updated successfully"})
	return nil
}

// DeleteInvoiceItem godoc
// @Summary Remove um item de nota fiscal
// @Description Deleta um item de nota fiscal pelo ID
// @Tags invoice_item
// @Accept  json
// @Produce  json
// @Param id path int true "ID do item da nota"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/invoice_items/{id} [delete]
func (r *InvoiceItemRepository) DeleteInvoiceItem(c *fiber.Ctx) error {
	id := c.Params("id")
	invoiceItem := models.InvoiceItem{}

	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id can not be empty"})
	}

	err := r.DB.Where("id = ?", id).First(&invoiceItem).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "invoice item not found"})
	}

	err = r.DB.Delete(&invoiceItem).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not delete invoice item"})
	}

	return c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "invoice item deleted successfully"})
}

func (r *InvoiceItemRepository) SetupInvoiceItemRoutes(app *fiber.App) {
	api := app.Group("/api", middleware.AuthRequired)
	api.Post("/invoice_items", r.CreateInvoiceItem)
	api.Get("/invoice_items", r.GetInvoiceItems)
	api.Get("/invoice_items/:id", r.GetInvoiceItemByID)
	api.Put("/invoice_items/:id", r.UpdateInvoiceItem)
	api.Delete("/invoice_items/:id", r.DeleteInvoiceItem)
}
