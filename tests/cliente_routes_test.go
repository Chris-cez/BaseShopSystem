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

func setupClientTestApp() (*fiber.App, *gorm.DB) {
	app := fiber.New()
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	models.MigrateClient(db)
	repo := routes.ClientRepository{DB: db}
	repo.SetupClientRoutes(app)
	return app, db
}

func TestCreateClient(t *testing.T) {
	app, db := setupClientTestApp()
	client := models.Client{
		Name:      "Cliente Teste",
		CPF:       "12345678900",
		AddressID: 1,
	}
	body, _ := json.Marshal(client)
	req := httptest.NewRequest("POST", "/api/clients", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
	var count int64
	db.Model(&models.Client{}).Where("cpf = ?", "12345678900").Count(&count)
	if count != 1 {
		t.Errorf("Cliente não foi criado no banco")
	}
}

func TestGetClients(t *testing.T) {
	app, db := setupClientTestApp()
	db.Create(&models.Client{Name: "Cliente Lista", CPF: "11122233344", AddressID: 1})
	req := httptest.NewRequest("GET", "/api/clients", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
}

func TestGetClientByID(t *testing.T) {
	app, db := setupClientTestApp()
	client := models.Client{Name: "Cliente ID", CPF: "22233344455", AddressID: 1}
	db.Create(&client)
	url := fmt.Sprintf("/api/clients/%d", client.ID)
	req := httptest.NewRequest("GET", url, nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
}

func TestUpdateClient(t *testing.T) {
	app, db := setupClientTestApp()
	client := models.Client{Name: "Cliente Atualizar", CPF: "33344455566", AddressID: 1}
	db.Create(&client)
	client.Name = "Cliente Atualizado"
	body, _ := json.Marshal(client)
	url := fmt.Sprintf("/api/clients/%d", client.ID)
	req := httptest.NewRequest("PUT", url, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
	var updated models.Client
	db.First(&updated, client.ID)
	if updated.Name != "Cliente Atualizado" {
		t.Errorf("Cliente não foi atualizado")
	}
}

func TestDeleteClient(t *testing.T) {
	app, db := setupClientTestApp()
	client := models.Client{Name: "Cliente Deletar", CPF: "44455566677", AddressID: 1}
	db.Create(&client)
	url := fmt.Sprintf("/api/clients/%d", client.ID)
	req := httptest.NewRequest("DELETE", url, nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
	var count int64
	db.Model(&models.Client{}).Where("id = ?", client.ID).Count(&count)
	if count != 0 {
		t.Errorf("Cliente não foi deletado")
	}
}
