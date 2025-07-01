package routes

import (
	"net/http"

	"github.com/Chris-cez/BaseShopSystem/middleware"
	"github.com/Chris-cez/BaseShopSystem/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ClassRepository struct {
	DB *gorm.DB
}

// CreateClass godoc
// @Summary Cria uma nova classe
// @Description Adiciona uma nova classe ao banco de dados
// @Tags class
// @Accept  json
// @Produce  json
// @Param class body models.Class true "Dados da classe"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /api/classes [post]
func (r *ClassRepository) CreateClass(c *fiber.Ctx) error {
	class := models.Class{}

	err := c.BodyParser(&class)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
		return err
	}
	err = r.DB.Create(&class).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create class"})
		return err
	}

	c.Status(http.StatusOK).JSON(
		&fiber.Map{"class": "class has been added"})
	return nil
}

// GetClasses godoc
// @Summary Lista todas as classes
// @Description Retorna todas as classes cadastradas
// @Tags class
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/classes [get]
func (r *ClassRepository) GetClasses(c *fiber.Ctx) error {
	classModels := []models.Class{}
	err := r.DB.Find(&classModels).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get classes"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "classes fetched successfully",
			"data": classModels})
	return nil
}

// GetClassByID godoc
// @Summary Busca classe por ID
// @Description Retorna uma classe pelo seu ID
// @Tags class
// @Accept  json
// @Produce  json
// @Param id path int true "ID da classe"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/classes/{id} [get]
func (r *ClassRepository) GetClassByID(c *fiber.Ctx) error {
	id := c.Params("id")
	classModel := models.Class{}

	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id can not be empty"})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(&classModel).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the class"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "class fetched successfully",
			"data": classModel})
	return nil
}

// UpdateClass godoc
// @Summary Atualiza uma classe
// @Description Atualiza os dados de uma classe existente
// @Tags class
// @Accept  json
// @Produce  json
// @Param id path int true "ID da classe"
// @Param class body models.Class true "Dados da classe"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/classes/{id} [put]
func (r *ClassRepository) UpdateClass(c *fiber.Ctx) error {
	id := c.Params("id")
	class := models.Class{}

	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id can not be empty"})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(&class).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "class not found"})
		return err
	}

	err = c.BodyParser(&class)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
		return err
	}

	err = r.DB.Save(&class).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not update class"})
		return err
	}

	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "class updated successfully"})
	return nil
}

func (r *ClassRepository) SetupClassRoutes(app *fiber.App) {
	api := app.Group("/api", middleware.AuthRequired)
	api.Post("/classes", r.CreateClass)
	api.Get("/classes", r.GetClasses)
	api.Get("/classes/:id", r.GetClassByID)
	api.Put("/classes/:id", r.UpdateClass)
}
