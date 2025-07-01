package routes

import (
	"net/http"

	"github.com/Chris-cez/BaseShopSystem/middleware"
	"github.com/Chris-cez/BaseShopSystem/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SaleRepository struct {
	DB *gorm.DB
}

// CreateDraftInvoice godoc
// @Summary Cria uma nota fiscal rascunho
// @Description Cria uma nota fiscal sem cliente e método de pagamento
// @Tags sale
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/sale/draft [post]
func (r *SaleRepository) CreateDraftInvoice(c *fiber.Ctx) error {
	invoice := models.Invoice{
		Numero: uuid.NewString(), // Gera um número único
	}
	if err := r.DB.Create(&invoice).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not create draft invoice"})
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{"invoice_id": invoice.Numero})
}

// AddItemToInvoice godoc
// @Summary Adiciona um item à nota fiscal
// @Description Adiciona um produto à nota fiscal informando quantidade
// @Tags sale
// @Accept  json
// @Produce  json
// @Param data body object true "Dados do item"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /api/sale/add_item [post]
func (r *SaleRepository) AddItemToInvoice(c *fiber.Ctx) error {
	type AddItemRequest struct {
		InvoiceID string `json:"invoice_id"`
		ProductID uint   `json:"product_id"`
		Quantity  int    `json:"quantity"`
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

// FinalizeInvoice godoc
// @Summary Finaliza a nota fiscal
// @Description Finaliza a nota fiscal informando cliente e método de pagamento
// @Tags sale
// @Accept  json
// @Produce  json
// @Param data body object true "Dados para finalizar nota"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /api/sale/finalize [post]
func (r *SaleRepository) FinalizeInvoice(c *fiber.Ctx) error {
	type FinalizeRequest struct {
		InvoiceID       string `json:"invoice_id"`
		ClientID        uint   `json:"client_id"`
		PaymentMethodID uint   `json:"payment_method_id"`
	}
	var req FinalizeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "invalid request"})
	}
	var invoice models.Invoice
	if err := r.DB.First(&invoice, "numero = ?", req.InvoiceID).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "invoice not found"})
	}
	invoice.ClientID = req.ClientID
	invoice.PaymentMethodID = req.PaymentMethodID
	if err := r.DB.Save(&invoice).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not finalize invoice"})
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{"message": "invoice finalized"})
}

// GetInvoiceItems godoc
// @Summary Consulta itens de uma nota fiscal
// @Description Retorna todos os itens de uma nota fiscal pelo ID
// @Tags sale
// @Accept  json
// @Produce  json
// @Param invoice_id path string true "ID da nota fiscal"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/sale/items/{invoice_id} [get]
func (r *SaleRepository) GetInvoiceItems(c *fiber.Ctx) error {
	invoiceID := c.Params("invoice_id")
	var items []models.InvoiceItem
	if err := r.DB.Where("invoice_id = ?", invoiceID).Find(&items).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get items"})
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{"items": items})
}

func (r *SaleRepository) SetupSaleRoutes(app *fiber.App) {
	api := app.Group("/api/sale", middleware.AuthRequired)
	api.Post("/draft", r.CreateDraftInvoice)
	api.Post("/add_item", r.AddItemToInvoice)
	api.Post("/finalize", r.FinalizeInvoice)
	api.Get("/items/:invoice_id", r.GetInvoiceItems)
}
