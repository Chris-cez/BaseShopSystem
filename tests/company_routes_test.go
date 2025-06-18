package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Chris-cez/BaseShopSystem/models"
	"github.com/Chris-cez/BaseShopSystem/routes"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupCompanyTestApp() (*fiber.App, *gorm.DB) {
	app := fiber.New()
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	models.MigrateCompany(db)
	repo := routes.CompanyRepository{DB: db}
	repo.SetupCompanyRoutes(app)
	return app, db
}

func createTestCompany(db *gorm.DB) models.Company {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("senha123"), bcrypt.DefaultCost)
	company := models.Company{
		Name:              "Empresa Teste",
		CNPJ:              "12345678000199",
		InscricaoEstadual: "ISENTO",
		Password:          string(hashedPassword),
		Address_id:        1,
	}
	db.Create(&company)
	return company
}

func TestCreateCompany(t *testing.T) {
	app, db := setupCompanyTestApp()
	company := models.Company{
		Name:              "Empresa Nova",
		CNPJ:              "98765432000199",
		InscricaoEstadual: "ISENTO",
		Password:          "senha123",
		Address_id:        1,
	}
	body, _ := json.Marshal(company)
	req := httptest.NewRequest("POST", "/api/company", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
	var count int64
	db.Model(&models.Company{}).Where("cnpj = ?", "98765432000199").Count(&count)
	if count != 1 {
		t.Errorf("Empresa não foi criada no banco")
	}
}

func TestAuthenticateCompany(t *testing.T) {
	app, db := setupCompanyTestApp()
	company := createTestCompany(db)
	authReq := map[string]string{
		"cnpj":     company.CNPJ,
		"password": "senha123",
	}
	body, _ := json.Marshal(authReq)
	req := httptest.NewRequest("POST", "/api/entrar", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	if result["token"] == nil {
		t.Errorf("Token não foi retornado na autenticação")
	}
}

func TestGetCompanies(t *testing.T) {
	app, db := setupCompanyTestApp()
	createTestCompany(db)
	req := httptest.NewRequest("GET", "/api/company", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
}

func TestGetCompanyByID(t *testing.T) {
	app, db := setupCompanyTestApp()
	company := createTestCompany(db)
	url := fmt.Sprintf("/api/company/%d", company.ID)
	req := httptest.NewRequest("GET", url, nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
}

func TestUpdateCompany(t *testing.T) {
	app, db := setupCompanyTestApp()
	company := createTestCompany(db)
	company.Name = "Empresa Atualizada"
	body, _ := json.Marshal(company)
	url := fmt.Sprintf("/api/company/%d", company.ID)
	req := httptest.NewRequest("PUT", url, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
	var updated models.Company
	db.First(&updated, company.ID)
	if updated.Name != "Empresa Atualizada" {
		t.Errorf("Empresa não foi atualizada")
	}
}
