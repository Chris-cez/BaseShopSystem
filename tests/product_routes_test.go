package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Chris-cez/BaseShopSystem/middleware"
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
		Code:        "001",
		Price:       10.0,
		Name:        "Produto Teste",
		NCM:         "12345678",
		GTIN:        "7891234567890",
		UM:          "UN",
		Description: "Descrição teste",
		ClassID:     1,
		Stock:       100,
		ValTrib:     0.5,
	}
	body, _ := json.Marshal(product)
	req := httptest.NewRequest("POST", "/api/products", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+getTestProductToken())
	resp, err := app.Test(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("Erro ao criar produto: %v", err)
	}
	var createdProduct models.Product
	db.Find(&createdProduct, "code = ?", "001")
	return createdProduct.ID
}

func getTestProductToken() string {
	token, _ := middleware.GenerateJWT("12345678000199")
	return token
}

func TestCreateProduct(t *testing.T) {
	app, db := setupTestApp()

	product := map[string]interface{}{
		"code":        "001",
		"price":       10.0,
		"name":        "Produto Teste",
		"ncm":         "12345678",
		"gtin":        "7891234567890",
		"um":          "UN",
		"description": "Descrição teste",
		"class_id":    1,
		"stock":       100,
		"valtrib":     0.5,
	}
	body, _ := json.Marshal(product)
	req := httptest.NewRequest("POST", "/api/products", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+getTestProductToken())
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}

	var count int64
	db.Model(&models.Product{}).Where("code = ?", "001").Count(&count)
	if count != 1 {
		t.Errorf("Produto não foi criado no banco")
	}
}

func TestGetProducts(t *testing.T) {
	app, db := setupTestApp()
	db.Create(&models.Product{
		Code:        "003",
		Price:       30.0,
		Name:        "Produto Lista",
		NCM:         "11111111",
		GTIN:        "7891111111111",
		UM:          "UN",
		Description: "Descrição lista",
		ClassID:     3,
		Stock:       10,
		ValTrib:     0.2,
	})
	req := httptest.NewRequest("GET", "/api/products", nil)
	req.Header.Set("Authorization", "Bearer "+getTestProductToken())
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
	prod := models.Product{
		Code:        "004",
		Price:       40.0,
		Name:        "Produto ID",
		NCM:         "22222222",
		GTIN:        "7892222222222",
		UM:          "UN",
		Description: "Descrição id",
		ClassID:     4,
		Stock:       20,
		ValTrib:     0.3,
	}
	db.Create(&prod)
	url := fmt.Sprintf("/api/products/%d", prod.ID)
	req := httptest.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+getTestProductToken())
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
	db.Create(&models.Product{
		Code:        "005",
		Price:       50.0,
		Name:        "Produto Nome",
		NCM:         "33333333",
		GTIN:        "7893333333333",
		UM:          "UN",
		Description: "Descrição nome",
		ClassID:     5,
		Stock:       30,
		ValTrib:     0.4,
	})
	req := httptest.NewRequest("GET", "/api/products/name/Produto", nil)
	req.Header.Set("Authorization", "Bearer "+getTestProductToken())
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
	db.Create(&models.Product{
		Code:        "006",
		Price:       60.0,
		Name:        "Produto Classe",
		NCM:         "44444444",
		GTIN:        "7894444444444",
		UM:          "UN",
		Description: "Descrição classe",
		ClassID:     123,
		Stock:       40,
		ValTrib:     0.6,
	})
	req := httptest.NewRequest("GET", "/api/products/class/123", nil)
	req.Header.Set("Authorization", "Bearer "+getTestProductToken())
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

	prod := models.Product{
		Code:        "002",
		Price:       20.0,
		Name:        "Produto Atualizar",
		NCM:         "87654321",
		GTIN:        "7890987654321",
		UM:          "UN",
		Description: "Descrição atualizar",
		ClassID:     2,
		Stock:       50,
		ValTrib:     1.0,
	}
	db.Create(&prod)

	update := map[string]interface{}{
		"code":        "002",
		"price":       25.0,
		"name":        "Produto Atualizado",
		"ncm":         "87654321",
		"gtin":        "7890987654321",
		"um":          "UN",
		"description": "Descrição atualizada",
		"class_id":    2,
		"stock":       60,
		"valtrib":     1.5,
	}
	body, _ := json.Marshal(update)
	url := fmt.Sprintf("/api/products/%d", prod.ID)
	req := httptest.NewRequest("PUT", url, bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+getTestProductToken())
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
	if updated.Name != "Produto Atualizado" || updated.Price != 25.0 || updated.Stock != 60 || updated.ValTrib != 1.5 {
		t.Errorf("Produto não foi atualizado corretamente")
	}
}

func TestDeleteProduct(t *testing.T) {
	app, db := setupTestApp()
	prod := models.Product{
		Code:        "008",
		Price:       80.0,
		Name:        "Produto Deletar",
		NCM:         "66666666",
		GTIN:        "7896666666666",
		UM:          "UN",
		Description: "Descrição deletar",
		ClassID:     8,
		Stock:       80,
		ValTrib:     0.8,
	}
	db.Create(&prod)
	url := fmt.Sprintf("/api/products/%d", prod.ID)
	req := httptest.NewRequest("DELETE", url, nil)
	req.Header.Set("Authorization", "Bearer "+getTestProductToken())
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
