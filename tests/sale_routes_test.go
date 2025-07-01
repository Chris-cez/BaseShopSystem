package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Chris-cez/BaseShopSystem/models"
	"github.com/Chris-cez/BaseShopSystem/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Função auxiliar para gerar um token JWT válido para os testes de sale
func getTestSaleToken() string {
	claims := jwt.MapClaims{
		"cnpj": "00000000000191",
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := []byte("ed45w3ecrtdxs2t1nmftvby")
	tokenString, _ := token.SignedString(jwtSecret)
	return tokenString
}

func setupSaleTestApp() (*fiber.App, *gorm.DB) {
	app := fiber.New()
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	models.MigrateInvoice(db)
	models.MigrateInvoiceItem(db)
	models.MigrateClient(db)
	models.MigratePaymentMethod(db)
	models.MigrateProduct(db)

	db.Create(&models.Client{Name: "Cliente Venda", CPF: "99988877766", AddressID: 1})
	db.Create(&models.PaymentMethod{Name: "dinheiro fisico"})
	db.Create(&models.Product{Code: "P001", Price: 5.0, Name: "Produto Venda"})

	repo := routes.SaleRepository{DB: db}
	repo.SetupSaleRoutes(app)
	return app, db
}

func TestCreateDraftInvoice(t *testing.T) {
	app, db := setupSaleTestApp()
	req := httptest.NewRequest("POST", "/api/sale/draft", nil)
	req.Header.Set("Authorization", "Bearer "+getTestSaleToken())
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
	var count int64
	db.Model(&models.Invoice{}).Count(&count)
	if count == 0 {
		t.Errorf("Draft invoice não foi criada")
	}
}

func TestAddItemToInvoice(t *testing.T) {
	app, db := setupSaleTestApp()
	invoice := models.Invoice{Numero: uuid.NewString()}
	db.Create(&invoice)
	// Cria um produto
	product := models.Product{Code: "P002", Price: 10.0, Name: "Produto 2"}
	db.Create(&product)
	// Adiciona item à nota
	itemReq := map[string]interface{}{
		"invoice_id": invoice.Numero,
		"product_id": product.ID,
		"quantity":   2,
	}
	body, _ := json.Marshal(itemReq)
	req := httptest.NewRequest("POST", "/api/sale/add_item", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+getTestSaleToken())
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
	var count int64
	db.Model(&models.InvoiceItem{}).Where("invoice_id = ?", invoice.Numero).Count(&count)
	if count == 0 {
		t.Errorf("Item não foi adicionado à nota")
	}
}

func TestFinalizeInvoice(t *testing.T) {
	app, db := setupSaleTestApp()
	// Cria nota e cliente
	client := models.Client{Name: "Cliente Finalizar", CPF: "88877766655", AddressID: 1}
	db.Create(&client)
	payment := models.PaymentMethod{Name: "cartao"}
	db.Create(&payment)
	invoice := models.Invoice{Numero: uuid.NewString()}
	db.Create(&invoice)
	// Finaliza nota
	finalizeReq := map[string]interface{}{
		"invoice_id":        invoice.Numero,
		"client_id":         client.ID,
		"payment_method_id": payment.ID,
	}
	body, _ := json.Marshal(finalizeReq)
	req := httptest.NewRequest("POST", "/api/sale/finalize", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+getTestSaleToken())
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
	var updated models.Invoice
	db.First(&updated, "numero = ?", invoice.Numero)
	if updated.ClientID != client.ID || updated.PaymentMethodID != payment.ID {
		t.Errorf("Invoice não foi finalizada corretamente")
	}
}

func TestGetInvoiceItems(t *testing.T) {
	app, db := setupSaleTestApp()
	// Cria nota e item
	invoice := models.Invoice{Numero: uuid.NewString()}
	db.Create(&invoice)
	product := models.Product{Code: "P003", Price: 15.0, Name: "Produto 3"}
	db.Create(&product)
	item := models.InvoiceItem{InvoiceID: invoice.Numero, ProductID: product.ID, Quantity: 1}
	db.Create(&item)
	url := fmt.Sprintf("/api/sale/items/%s", invoice.Numero)
	req := httptest.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+getTestSaleToken())
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
}
