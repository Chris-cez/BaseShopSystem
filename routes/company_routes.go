package routes

import (
	"net/http"

	"github.com/Chris-cez/BaseShopSystem/middleware"
	"github.com/Chris-cez/BaseShopSystem/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CompanyRepository struct {
	DB *gorm.DB
}

// CreateCompany godoc
// @Summary Cria uma nova empresa
// @Description Adiciona uma nova empresa ao banco de dados
// @Tags company
// @Accept  json
// @Produce  json
// @Param company body models.Company true "Dados da empresa"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /api/company [post]
func (r *CompanyRepository) CreateCompany(c *fiber.Ctx) error {
	company := models.Company{}

	err := c.BodyParser(&company)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(company.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "Could not hash password"})
		return err
	}
	company.Password = string(hashedPassword)

	err = r.DB.Create(&company).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create company"})
		return err
	}

	c.Status(http.StatusOK).JSON(
		&fiber.Map{"company": "company has been added"})
	return nil
}

// GetCompanies godoc
// @Summary Lista todas as empresas
// @Description Retorna todas as empresas cadastradas
// @Tags company
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/company [get]
func (r *CompanyRepository) GetCompanies(c *fiber.Ctx) error {
	companyModels := []models.Company{}
	err := r.DB.Find(&companyModels).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get companies"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "companies fetched successfully",
			"data": companyModels})
	return nil
}

// GetCompanyByID godoc
// @Summary Busca empresa por ID
// @Description Retorna uma empresa pelo seu ID
// @Tags company
// @Accept  json
// @Produce  json
// @Param id path int true "ID da empresa"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/company/{id} [get]
func (r *CompanyRepository) GetCompanyByID(c *fiber.Ctx) error {
	id := c.Params("id")
	companyModel := models.Company{}

	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id can not be empty"})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(&companyModel).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get company"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "company fetched successfully",
			"data": companyModel})
	return nil
}

// UpdateCompany godoc
// @Summary Atualiza uma empresa
// @Description Atualiza os dados de uma empresa existente
// @Tags company
// @Accept  json
// @Produce  json
// @Param id path int true "ID da empresa"
// @Param company body models.Company true "Dados da empresa"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/company/{id} [put]
func (r *CompanyRepository) UpdateCompany(c *fiber.Ctx) error {
	id := c.Params("id")
	company := models.Company{}

	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id can not be empty"})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(&company).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get company"})
		return err
	}

	err = c.BodyParser(&company)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
		return err
	}
	err = r.DB.Save(&company).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not update company"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "company updated successfully",
			"data": company})
	return nil
}

// AuthenticateCompany godoc
// @Summary Autentica uma empresa
// @Description Autentica uma empresa pelo CNPJ e senha, retornando um token JWT
// @Tags company
// @Accept  json
// @Produce  json
// @Param credentials body object true "CNPJ e senha"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/entrar [post]
func (r *CompanyRepository) AuthenticateCompany(c *fiber.Ctx) error {
	type AuthRequest struct {
		CNPJ     string `json:"cnpj"`
		Password string `json:"password"`
	}

	var authRequest AuthRequest
	if err := c.BodyParser(&authRequest); err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Invalid request body"})
		return err
	}

	// Buscar a empresa pelo CNPJ
	company := models.Company{}
	err := r.DB.Where("cnpj = ?", authRequest.CNPJ).First(&company).Error
	if err != nil {
		c.Status(http.StatusUnauthorized).JSON(
			&fiber.Map{"message": "Invalid CNPJ or password"})
		return err
	}

	// Verificar a senha (comparar o hash)
	err = bcrypt.CompareHashAndPassword([]byte(company.Password), []byte(authRequest.Password))
	if err != nil {
		c.Status(http.StatusUnauthorized).JSON(
			&fiber.Map{"message": "Invalid CNPJ or password"})
		return nil
	}

	// Gerar o token JWT
	token, err := middleware.GenerateJWT(company.CNPJ)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "Could not generate token"})
		return err
	}

	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Authentication successful", "token": token})
	return nil
}

func (r *CompanyRepository) SetupCompanyRoutes(app *fiber.App) {
	api := app.Group("/api", middleware.AuthRequired)

	api.Post("/company", r.CreateCompany)
	api.Get("/company", r.GetCompanies)
	api.Get("/company/:id", r.GetCompanyByID)
	api.Put("/company/:id", r.UpdateCompany)
	app.Post("/entrar", r.AuthenticateCompany) // Nova rota de autenticação
}
