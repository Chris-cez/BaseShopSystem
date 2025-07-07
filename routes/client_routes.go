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

// CreateClient godoc
// @Summary Cria um novo cliente
// @Description Adiciona um novo cliente ao banco de dados
// @Tags client
// @Accept  json
// @Produce  json
// @Param client body models.Client true "Dados do cliente"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /api/clients [post]
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

// GetClients godoc
// @Summary Lista todos os clientes
// @Description Retorna todos os clientes cadastrados
// @Tags client
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/clients [get]
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

// GetClientByID godoc
// @Summary Busca cliente por ID
// @Description Retorna um cliente pelo seu ID
// @Tags client
// @Accept  json
// @Produce  json
// @Param id path int true "ID do cliente"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/clients/{id} [get]
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

// UpdateClient godoc
// @Summary Atualiza um cliente
// @Description Atualiza os dados de um cliente existente
// @Tags client
// @Accept  json
// @Produce  json
// @Param id path int true "ID do cliente"
// @Param client body models.Client true "Dados do cliente"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/clients/{id} [put]
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

// DeleteClient godoc
// @Summary Remove um cliente
// @Description Deleta um cliente pelo ID
// @Tags client
// @Accept  json
// @Produce  json
// @Param id path int true "ID do cliente"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/clients/{id} [delete]
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

// GetClientsByName godoc
// @Summary Busca clientes por nome
// @Description Retorna clientes que contenham o nome informado
// @Tags client
// @Accept  json
// @Produce  json
// @Param name path string true "Nome do cliente"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/clients/name/{name} [get]
func (r *ClientRepository) GetClientsByName(c *fiber.Ctx) error {
	name := c.Params("name")
	clientModels := []models.Client{}

	err := r.DB.Where("name ILIKE ?", "%"+name+"%").Find(&clientModels).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get clients by name"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "clients fetched successfully",
			"data": clientModels})
	return nil
}

// GetClientsByCPF godoc
// @Summary Busca cliente por CPF
// @Description Retorna um cliente pelo CPF informado
// @Tags client
// @Accept  json
// @Produce  json
// @Param cpf path string true "CPF do cliente"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/clients/cpf/{cpf} [get]
func (r *ClientRepository) GetClientsByCPF(c *fiber.Ctx) error {
	cpf := c.Params("cpf")
	clientModel := models.Client{}

	err := r.DB.Where("cpf = ?", cpf).First(&clientModel).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get client by CPF"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "client fetched successfully",
			"data": clientModel})
	return nil
}

func (r *ClientRepository) SetupClientRoutes(app *fiber.App) {
	api := app.Group("/api", middleware.AuthRequired)
	api.Post("/clients", r.CreateClient)
	api.Get("/clients", r.GetClients)
	api.Get("/clients/name/:name", r.GetClientsByName)
	api.Get("/clients/cpf/:cpf", r.GetClientsByCPF)
	api.Get("/clients/:id", r.GetClientByID)
	api.Put("/clients/:id", r.UpdateClient)
	api.Delete("/clients/:id", r.DeleteClient)
}
