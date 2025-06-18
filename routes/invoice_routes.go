package routes

import (
    "net/http"

    "github.com/Chris-cez/BaseShopSystem/models"
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

type InvoiceRepository struct {
    DB *gorm.DB
}

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
    api := app.Group("/api")
    api.Post("/invoices", r.CreateInvoice)
    api.Get("/invoices", r.GetInvoices)
    api.Get("/invoices/:id", r.GetInvoiceByID)
}