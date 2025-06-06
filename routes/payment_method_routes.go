package routes

import (
	"net/http"

	"github.com/Chris-cez/BaseShopSystem/middleware"
	"github.com/Chris-cez/BaseShopSystem/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PaymentMethodRepository struct {
	DB *gorm.DB
}

func (r *PaymentMethodRepository) CreatePaymentMethod(c *fiber.Ctx) error {
	paymentMethod := models.PaymentMethod{}

	err := c.BodyParser(&paymentMethod)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
		return err
	}
	err = r.DB.Create(&paymentMethod).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create payment method"})
		return err
	}

	c.Status(http.StatusOK).JSON(
		&fiber.Map{"payment_method": "payment method has been added"})
	return nil
}

func (r *PaymentMethodRepository) GetPaymentMethods(c *fiber.Ctx) error {
	paymentMethodModels := []models.PaymentMethod{}
	err := r.DB.Find(&paymentMethodModels).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get payment methods"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "payment methods fetched successfully",
			"data": paymentMethodModels})
	return nil
}

func (r *PaymentMethodRepository) GetPaymentMethodByID(c *fiber.Ctx) error {
	id := c.Params("id")
	paymentMethodModel := models.PaymentMethod{}

	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id can not be empty"})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(&paymentMethodModel).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the payment method"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "payment method fetched successfully",
			"data": paymentMethodModel})
	return nil
}

func (r *PaymentMethodRepository) DeletePaymentMethod(c *fiber.Ctx) error {
	id := c.Params("id")
	paymentMethod := models.PaymentMethod{}

	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id can not be empty"})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(&paymentMethod).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "payment method not found"})
		return err
	}

	err = r.DB.Delete(&paymentMethod).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not delete payment method"})
		return err
	}

	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "payment method deleted successfully"})
	return nil
}

func (r *PaymentMethodRepository) SetupPaymentMethodRoutes(app *fiber.App) {
	api := app.Group("/api", middleware.AuthRequired)
	api.Post("/payment_methods", r.CreatePaymentMethod)
	api.Get("/payment_methods", r.GetPaymentMethods)
	api.Get("/payment_methods/:id", r.GetPaymentMethodByID)
	api.Delete("/payment_methods/:id", r.DeletePaymentMethod)
}
