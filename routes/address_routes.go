package routes

import (
	"net/http"

	"github.com/Chris-cez/BaseShopSystem/middleware"
	"github.com/Chris-cez/BaseShopSystem/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AddressRepository struct {
	DB *gorm.DB
}

// CreateAddress godoc
// @Summary Cria um novo endereço
// @Description Adiciona um novo endereço ao banco de dados
// @Tags address
// @Accept  json
// @Produce  json
// @Param address body models.Address true "Dados do endereço"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /api/addresses [post]
func (r *AddressRepository) CreateAddress(c *fiber.Ctx) error {
	address := models.Address{}

	err := c.BodyParser(&address)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
		return err
	}
	err = r.DB.Create(&address).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create address"})
		return err
	}

	c.Status(http.StatusOK).JSON(
		&fiber.Map{"address": "address has been added"})
	return nil
}

// GetAddresses godoc
// @Summary Lista todos os endereços
// @Description Retorna todos os endereços cadastrados
// @Tags address
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/addresses [get]
func (r *AddressRepository) GetAddresses(c *fiber.Ctx) error {
	addressModels := []models.Address{}
	err := r.DB.Find(&addressModels).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get addresses"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "addresses fetched successfully",
			"data": addressModels})
	return nil
}

// GetAddressByID godoc
// @Summary Busca endereço por ID
// @Description Retorna um endereço pelo seu ID
// @Tags address
// @Accept  json
// @Produce  json
// @Param id path int true "ID do endereço"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/addresses/{id} [get]
func (r *AddressRepository) GetAddressByID(c *fiber.Ctx) error {
	id := c.Params("id")
	addressModel := models.Address{}

	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id can not be empty"})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(&addressModel).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the address"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "address fetched successfully",
			"data": addressModel})
	return nil
}

// UpdateAddress godoc
// @Summary Atualiza um endereço
// @Description Atualiza os dados de um endereço existente
// @Tags address
// @Accept  json
// @Produce  json
// @Param id path int true "ID do endereço"
// @Param address body models.Address true "Dados do endereço"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/addresses/{id} [put]
func (r *AddressRepository) UpdateAddress(c *fiber.Ctx) error {
	id := c.Params("id")
	address := models.Address{}

	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id can not be empty"})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(&address).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "address not found"})
		return err
	}

	err = c.BodyParser(&address)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
		return err
	}

	err = r.DB.Save(&address).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not update address"})
		return err
	}

	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "address updated successfully"})
	return nil
}

func (r *AddressRepository) SetupAddressRoutes(app *fiber.App) {
	api := app.Group("/api", middleware.AuthRequired)
	api.Post("/addresses", r.CreateAddress)
	api.Get("/addresses", r.GetAddresses)
	api.Get("/addresses/:id", r.GetAddressByID)
	api.Put("/addresses/:id", r.UpdateAddress)
}
