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

func (r *InvoiceItemRepository) CreateInvoiceItem(c *fiber.Ctx) error {
	invoiceItem := models.InvoiceItem{}

	err := c.BodyParser(&invoiceItem)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
		return err
	}
	err = r.DB.Create(&invoiceItem).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create invoice item"})
		return err
	}

	c.Status(http.StatusOK).JSON(
		&fiber.Map{"invoice_item": "invoice item has been added"})
	return nil
}

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

func (r *InvoiceItemRepository) SetupInvoiceItemRoutes(app *fiber.App) {
	api := app.Group("/api", middleware.AuthRequired)
	api.Post("/invoice_items", r.CreateInvoiceItem)
	api.Get("/invoice_items", r.GetInvoiceItems)
	api.Get("/invoice_items/:id", r.GetInvoiceItemByID)
	api.Put("/invoice_items/:id", r.UpdateInvoiceItem)
}
