package main

import (
	"github.com/Chris-cez/BaseShopSystem/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func InsertTestData(db *gorm.DB) error {
	// Verifica se já existe o endereço de teste
	var count int64
	db.Model(&models.Address{}).Where("cep = ?", "12345678").Count(&count)
	if count == 0 {
		address := models.Address{
			Logradouro:  "Rua Teste",
			Numero:      "123",
			Complemento: "Apto 1",
			Bairro:      "Centro",
			Municipio:   "Cidade Teste",
			UF:          "SP",
			CEP:         "12345678",
		}
		if err := db.Create(&address).Error; err != nil {
			return err
		}

		// Cria empresa de teste associada ao endereço
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("senha123"), bcrypt.DefaultCost)
		company := models.Company{
			Name:              "Empresa Teste",
			CNPJ:              "12345678000199",
			InscricaoEstadual: "ISENTO",
			Password:          string(hashedPassword),
			Address_id:        int(address.ID),
		}
		if err := db.Create(&company).Error; err != nil {
			return err
		}
	}

	// Popula métodos de pagamento se não existirem
	var pmCount int64
	db.Model(&models.PaymentMethod{}).Count(&pmCount)
	if pmCount == 0 {
		paymentMethods := []models.PaymentMethod{
			{Name: "dinheiro fisico"},
			{Name: "cartão de crédito"},
			{Name: "cartão de débito"},
			{Name: "vale alimentação"},
			{Name: "pix"},
		}
		if err := db.Create(&paymentMethods).Error; err != nil {
			return err
		}
	}

	// Popula clientes de teste se não existirem
	var clientCount int64
	db.Model(&models.Client{}).Count(&clientCount)
	if clientCount == 0 {
		clients := []models.Client{
			{Name: "João da Silva", CPF: "12345678901", AddressID: 1},
			{Name: "Maria Oliveira", CPF: "98765432100", AddressID: 1},
			{Name: "Carlos Souza", CPF: "11122233344", AddressID: 1},
		}
		if err := db.Create(&clients).Error; err != nil {
			return err
		}
	}

	return nil
}
