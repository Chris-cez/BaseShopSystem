package routes

import (
	"github.com/Chris-cez/BaseShopSystem/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
)

type CompanyRepository struct {
	DB *gorm.DB
}

func (r *CompanyRepository) CreateCompany(c *fiber.Ctx) error {
	company := models.Company{}

	err := c.BodyParser(&company)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
		return err
	}
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
func (r *CompanyRepository) SetupCompanyRoutes(app *fiber.App) {
	api := app.Group("/api")
	company := api.Group("/company")

	company.Post("/", r.CreateCompany)
	company.Get("/", r.GetCompanies)
	company.Get("/:id", r.GetCompanyByID)
	company.Put("/:id", r.UpdateCompany)
}
