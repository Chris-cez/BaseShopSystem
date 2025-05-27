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
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestApp() (*fiber.App, *gorm.DB) {
	app := fiber.New()
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	models.MigrateProduct(db)
	repo := routes.ProductRepository{DB: db}
	repo.SetupProductRoutes(app)
	return app, db
}

func createProduct(app *fiber.App, t *testing.T, db *gorm.DB) uint {
	product := models.Product{
		Code:  "001",
		Price: 10.0,
		Name:  "Produto Teste",
	}
	body, _ := json.Marshal(product)
	req := httptest.NewRequest("POST", "/api/products", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("Erro ao criar produto: %v", err)
	}
	// Buscar o ID do produto criado
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	var createdProduct models.Product
	db.Find(&createdProduct, "code = ?", "001")
	return createdProduct.ID
}

func TestCreateProduct(t *testing.T) {
	app, db := setupTestApp()
	product := models.Product{
		Code:  "002",
		Price: 20.0,
		Name:  "Produto Novo",
	}
	body, _ := json.Marshal(product)
	req := httptest.NewRequest("POST", "/api/products", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
	var count int64
	db.Model(&models.Product{}).Where("code = ?", "002").Count(&count)
	if count != 1 {
		t.Errorf("Produto não foi criado no banco")
	}
}

func TestGetProducts(t *testing.T) {
	app, db := setupTestApp()
	// Cria produto para garantir que a lista não esteja vazia
	db.Create(&models.Product{Code: "003", Price: 30.0, Name: "Produto Lista"})
	req := httptest.NewRequest("GET", "/api/products", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
}

func TestGetProductByID(t *testing.T) {
	app, db := setupTestApp()
	prod := models.Product{Code: "004", Price: 40.0, Name: "Produto ID"}
	db.Create(&prod)
	url := fmt.Sprintf("/api/products/%d", prod.ID)
	req := httptest.NewRequest("GET", url, nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
}

func TestGetProductsByName(t *testing.T) {
	app, db := setupTestApp()
	db.Create(&models.Product{Code: "005", Price: 50.0, Name: "Produto Nome"})
	req := httptest.NewRequest("GET", "/api/products/name/Produto", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
}

func TestGetProductsByClass(t *testing.T) {
	app, db := setupTestApp()
	db.Create(&models.Product{Code: "006", Price: 60.0, Name: "Produto Classe", ClassID: 123})
	req := httptest.NewRequest("GET", "/api/products/class/123", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
}

func TestUpdateProduct(t *testing.T) {
	app, db := setupTestApp()
	prod := models.Product{Code: "007", Price: 70.0, Name: "Produto Atualizar"}
	db.Create(&prod)
	prod.Name = "Produto Atualizado"
	body, _ := json.Marshal(prod)
	url := fmt.Sprintf("/api/products/%d", prod.ID)
	req := httptest.NewRequest("PUT", url, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
	var updated models.Product
	db.First(&updated, prod.ID)
	if updated.Name != "Produto Atualizado" {
		t.Errorf("Produto não foi atualizado")
	}
}

func TestDeleteProduct(t *testing.T) {
	app, db := setupTestApp()
	prod := models.Product{Code: "008", Price: 80.0, Name: "Produto Deletar"}
	db.Create(&prod)
	url := fmt.Sprintf("/api/products/%d", prod.ID)
	req := httptest.NewRequest("DELETE", url, nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
	var count int64
	db.Model(&models.Product{}).Where("id = ?", prod.ID).Count(&count)
	if count != 0 {
		t.Errorf("Produto não foi deletado")
	}
}
