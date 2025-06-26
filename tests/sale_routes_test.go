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

func setupSaleTestApp() (*fiber.App, *gorm.DB) {
	app := fiber.New()
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	models.MigrateInvoice(db)
	models.MigrateInvoiceItem(db)
	models.MigrateClient(db)
	models.MigratePaymentMethod(db)
	models.MigrateProduct(db)

	// Cria dependências básicas
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
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
	// Verifica se uma invoice foi criada
	var count int64
	db.Model(&models.Invoice{}).Count(&count)
	if count != 1 {
		t.Errorf("Draft invoice não foi criada")
	}
}

func TestAddItemToInvoice(t *testing.T) {
	app, db := setupSaleTestApp()
	// Cria draft
	invoice := models.Invoice{}
	db.Create(&invoice)
	var product models.Product
	db.First(&product)
	body, _ := json.Marshal(map[string]interface{}{
		"invoice_id": invoice.Numero,
		"product_id": product.ID,
		"quantity":   2,
	})
	req := httptest.NewRequest("POST", "/api/sale/add_item", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
	// Verifica se item foi adicionado
	var count int64
	db.Model(&models.InvoiceItem{}).Count(&count)
	if count != 1 {
		t.Errorf("Item não foi adicionado à nota")
	}
}

func TestFinalizeInvoice(t *testing.T) {
	app, db := setupSaleTestApp()
	// Cria draft
	invoice := models.Invoice{}
	db.Create(&invoice)
	var client models.Client
	var pm models.PaymentMethod
	db.First(&client)
	db.First(&pm)
	body, _ := json.Marshal(map[string]interface{}{
		"invoice_id":        invoice.Numero,
		"client_id":         client.ID,
		"payment_method_id": pm.ID,
	})
	req := httptest.NewRequest("POST", "/api/sale/finalize", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
	// Verifica se invoice foi atualizada
	var updated models.Invoice
	db.First(&updated, "numero = ?", invoice.Numero)
	if updated.ClientID != client.ID || updated.PaymentMethodID != pm.ID {
		t.Errorf("Invoice não foi finalizada corretamente")
	}
}

func TestGetInvoiceItems(t *testing.T) {
	app, db := setupSaleTestApp()
	// Cria draft e item
	invoice := models.Invoice{}
	db.Create(&invoice)
	item := models.InvoiceItem{InvoiceID: invoice.Numero, ProductID: 1, Quantity: 1, Price: 5.0, ValorTotal: 5.0}
	db.Create(&item)
	url := fmt.Sprintf("/api/sale/items/%s", invoice.Numero)
	req := httptest.NewRequest("GET", url, nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
}
