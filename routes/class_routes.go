package routes

import (
	"net/http"

	"github.com/Chris-cez/BaseShopSystem/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ClassRepository struct {
	DB *gorm.DB
}

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
	api := app.Group("/api")
	api.Post("/classes", r.CreateClass)
	api.Get("/classes", r.GetClasses)
	api.Get("/classes/:id", r.GetClassByID)
	api.Put("/classes/:id", r.UpdateClass)
}
