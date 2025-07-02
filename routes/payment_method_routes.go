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

// GetPaymentMethods godoc
// @Summary Lista todos os métodos de pagamento
// @Description Retorna todos os métodos de pagamento cadastrados
// @Tags payment_method
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/payment_methods [get]
func (r *PaymentMethodRepository) GetPaymentMethods(c *fiber.Ctx) error {
	paymentMethods := []models.PaymentMethod{}
	err := r.DB.Find(&paymentMethods).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get payment methods"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "payment methods fetched successfully",
			"data": paymentMethods})
	return nil
}

func (r *PaymentMethodRepository) SetupPaymentMethodRoutes(app *fiber.App) {
	api := app.Group("/api", middleware.AuthRequired)
	api.Get("/payment_methods", r.GetPaymentMethods)
}
