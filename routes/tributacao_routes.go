package routes

import (
	"net/http"

	"github.com/Chris-cez/BaseShopSystem/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TributacaoRepository struct {
	DB *gorm.DB
}

func (r *TributacaoRepository) CreateTributacao(c *fiber.Ctx) error {
	tributacao := models.Tributacao{}

	err := c.BodyParser(&tributacao)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
		return err
	}
	err = r.DB.Create(&tributacao).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create tributacao"})
		return err
	}

	c.Status(http.StatusOK).JSON(
		&fiber.Map{"tributacao": "tributacao has been added"})
	return nil
}

func (r *TributacaoRepository) GetTributacoes(c *fiber.Ctx) error {
	tributacaoModels := []models.Tributacao{}
	err := r.DB.Find(&tributacaoModels).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get tributacoes"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "tributacoes fetched successfully",
			"data": tributacaoModels})
	return nil
}

func (r *TributacaoRepository) GetTributacaoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	tributacaoModel := models.Tributacao{}

	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id can not be empty"})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(&tributacaoModel).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the tributacao"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "tributacao fetched successfully",
			"data": tributacaoModel})
	return nil
}

func (r *TributacaoRepository) UpdateTributacao(c *fiber.Ctx) error {
	id := c.Params("id")
	tributacao := models.Tributacao{}

	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id can not be empty"})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(&tributacao).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "tributacao not found"})
		return err
	}

	err = c.BodyParser(&tributacao)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
		return err
	}

	err = r.DB.Save(&tributacao).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not update tributacao"})
		return err
	}

	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "tributacao updated successfully"})
	return nil
}

func (r *TributacaoRepository) SetupTributacaoRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/tributacoes", r.CreateTributacao)
	api.Get("/tributacoes", r.GetTributacoes)
	api.Get("/tributacoes/:id", r.GetTributacaoByID)
	api.Put("/tributacoes/:id", r.UpdateTributacao)
}
