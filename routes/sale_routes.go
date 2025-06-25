package routes

import (
	"net/http"
	"strconv"

	"github.com/Chris-cez/BaseShopSystem/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type SaleRepository struct {
	DB *gorm.DB
}

// Cria uma nota "inválida" (sem cliente e método de pagamento)
func (r *SaleRepository) CreateDraftInvoice(c *fiber.Ctx) error {
	invoice := models.Invoice{
		// Campos como ClientID e PaymentMethodID ficam nulos
	}
	if err := r.DB.Create(&invoice).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not create draft invoice"})
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{"invoice_id": invoice.Numero})
}

// Adiciona um item à nota (invoice)
func (r *SaleRepository) AddItemToInvoice(c *fiber.Ctx) error {
	type AddItemRequest struct {
		InvoiceID uint `json:"invoice_id"`
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}
	var req AddItemRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "invalid request"})
	}
	item := models.InvoiceItem{
		InvoiceID: req.InvoiceID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}
	if err := r.DB.Create(&item).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not add item"})
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{"message": "item added"})
}

// Finaliza a nota, informando cliente e método de pagamento
func (r *SaleRepository) FinalizeInvoice(c *fiber.Ctx) error {
	type FinalizeRequest struct {
		InvoiceID       uint `json:"invoice_id"`
		ClientID        uint `json:"client_id"`
		PaymentMethodID uint `json:"payment_method_id"`
	}
	var req FinalizeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "invalid request"})
	}
	var invoice models.Invoice
	if err := r.DB.First(&invoice, req.InvoiceID).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "invoice not found"})
	}
	invoice.ClientID = req.ClientID
	invoice.PaymentMethodID = req.PaymentMethodID
	if err := r.DB.Save(&invoice).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not finalize invoice"})
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{"message": "invoice finalized"})
}

// Consulta os itens de uma nota
func (r *SaleRepository) GetInvoiceItems(c *fiber.Ctx) error {
	invoiceID, err := strconv.Atoi(c.Params("invoice_id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "invalid invoice id"})
	}
	var items []models.InvoiceItem
	if err := r.DB.Where("invoice_id = ?", invoiceID).Find(&items).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get items"})
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{"items": items})
}

func (r *SaleRepository) SetupSaleRoutes(app *fiber.App) {
	api := app.Group("/api/sale")
	api.Post("/draft", r.CreateDraftInvoice)
	api.Post("/add_item", r.AddItemToInvoice)
	api.Post("/finalize", r.FinalizeInvoice)
	api.Get("/items/:invoice_id", r.GetInvoiceItems)
}
