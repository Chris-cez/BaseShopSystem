package routes

import (
	"net/http"

	"github.com/Chris-cez/BaseShopSystem/middleware"
	"github.com/Chris-cez/BaseShopSystem/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type InvoiceRepository struct {
	DB *gorm.DB
}

// CreateInvoice godoc
// @Summary Cria uma nova nota fiscal
// @Description Adiciona uma nova nota fiscal ao banco de dados
// @Tags invoice
// @Accept  json
// @Produce  json
// @Param invoice body models.Invoice true "Dados da nota fiscal"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /api/invoices [post]
func (r *InvoiceRepository) CreateInvoice(c *fiber.Ctx) error {
	invoice := models.Invoice{}

	err := c.BodyParser(&invoice)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
		return err
	}
	err = r.DB.Create(&invoice).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create invoice"})
		return err
	}

	c.Status(http.StatusOK).JSON(
		&fiber.Map{"invoice": "invoice has been added"})
	return nil
}

// GetInvoices godoc
// @Summary Lista todas as notas fiscais
// @Description Retorna todas as notas fiscais cadastradas
// @Tags invoice
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/invoices [get]
func (r *InvoiceRepository) GetInvoices(c *fiber.Ctx) error {
	invoiceModels := []models.Invoice{}
	err := r.DB.Find(&invoiceModels).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get invoices"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "invoices fetched successfully",
			"data": invoiceModels})
	return nil
}

// GetInvoiceByID godoc
// @Summary Busca nota fiscal por ID
// @Description Retorna uma nota fiscal pelo seu ID
// @Tags invoice
// @Accept  json
// @Produce  json
// @Param id path int true "ID da nota fiscal"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/invoices/{id} [get]
func (r *InvoiceRepository) GetInvoiceByID(c *fiber.Ctx) error {
	id := c.Params("id")
	invoiceModel := models.Invoice{}

	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id can not be empty"})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(&invoiceModel).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the invoice"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "invoice fetched successfully",
			"data": invoiceModel})
	return nil
}

func (r *InvoiceRepository) SetupInvoiceRoutes(app *fiber.App) {
	api := app.Group("/api", middleware.AuthRequired)
	api.Post("/invoices", r.CreateInvoice)
	api.Get("/invoices", r.GetInvoices)
	api.Get("/invoices/:id", r.GetInvoiceByID)
}
