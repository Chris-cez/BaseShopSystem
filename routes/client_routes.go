package routes

import (
	"net/http"

	"github.com/Chris-cez/BaseShopSystem/middleware"
	"github.com/Chris-cez/BaseShopSystem/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ClientRepository struct {
	DB *gorm.DB
}

func (r *ClientRepository) CreateClient(c *fiber.Ctx) error {
	client := models.Client{}

	err := c.BodyParser(&client)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
		return err
	}
	err = r.DB.Create(&client).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create client"})
		return err
	}

	c.Status(http.StatusOK).JSON(
		&fiber.Map{"client": "client has been added"})
	return nil
}

func (r *ClientRepository) GetClients(c *fiber.Ctx) error {
	clientModels := []models.Client{}
	err := r.DB.Find(&clientModels).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get clients"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "clients fetched successfully",
			"data": clientModels})
	return nil
}

func (r *ClientRepository) GetClientByID(c *fiber.Ctx) error {
	id := c.Params("id")
	clientModel := models.Client{}

	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id can not be empty"})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(&clientModel).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the client"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "client fetched successfully",
			"data": clientModel})
	return nil
}

func (r *ClientRepository) UpdateClient(c *fiber.Ctx) error {
	id := c.Params("id")
	client := models.Client{}

	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id can not be empty"})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(&client).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "client not found"})
		return err
	}

	err = c.BodyParser(&client)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
		return err
	}

	err = r.DB.Save(&client).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not update client"})
		return err
	}

	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "client updated successfully"})
	return nil
}

func (r *ClientRepository) DeleteClient(c *fiber.Ctx) error {
	id := c.Params("id")
	client := models.Client{}

	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id can not be empty"})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(&client).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "client not found"})
		return err
	}

	err = r.DB.Delete(&client).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not delete client"})
		return err
	}

	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "client deleted successfully"})
	return nil
}

func (r *ClientRepository) SetupClientRoutes(app *fiber.App) {
	api := app.Group("/api", middleware.AuthRequired)
	api.Post("/clients", r.CreateClient)
	api.Get("/clients", r.GetClients)
	api.Get("/clients/:id", r.GetClientByID)
	api.Put("/clients/:id", r.UpdateClient)
	api.Delete("/clients/:id", r.DeleteClient)
}
