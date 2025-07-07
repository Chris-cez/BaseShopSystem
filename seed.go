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

	// Após criar o endereço
	var address models.Address
	db.Where("cep = ?", "12345678").First(&address)
	addressID := address.ID

	// Popula clientes de teste se não existirem
	var clientCount int64
	db.Model(&models.Client{}).Count(&clientCount)
	if clientCount == 0 {
		clients := []models.Client{
			{Name: "João da Silva", CPF: "12345678901", AddressID: addressID},
			{Name: "Maria Oliveira", CPF: "98765432100", AddressID: addressID},
			{Name: "Carlos Souza", CPF: "11122233344", AddressID: addressID},
		}
		if err := db.Create(&clients).Error; err != nil {
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

	// Popula classes de produtos se não existirem
	var classCount int64
	db.Model(&models.Class{}).Count(&classCount)
	if classCount == 0 {
		classes := []models.Class{
			{Name: "Chocolate Meio Amargo", Description: "Produtos de chocolate meio amargo", NCM: "17049010"},
			{Name: "Chocolate Branco", Description: "Produtos de chocolate branco", NCM: "17049020"},
			{Name: "Bolacha", Description: "Bolachas e biscoitos", NCM: "19053100"},
			{Name: "Bombom", Description: "Bombons diversos", NCM: "17049030"},
			{Name: "Flores", Description: "Flores naturais e artificiais", NCM: "06039000"},
		}
		if err := db.Create(&classes).Error; err != nil {
			return err
		}
	}

	// Popula produtos de exemplo se não existirem
	var productCount int64
	db.Model(&models.Product{}).Count(&productCount)
	if productCount == 0 {
		// Busca os IDs das classes criadas
		var classes []models.Class
		db.Find(&classes)
		classMap := make(map[string]uint)
		for _, c := range classes {
			classMap[c.Name] = c.ID
		}

		products := []models.Product{
			{
				Code:        "CHMA01",
				Price:       12.50,
				Name:        "Chocolate Meio Amargo 70%",
				NCM:         "17049010",
				GTIN:        "7891000000011",
				UM:          "UN",
				Description: "Barra de chocolate meio amargo 70% cacau",
				ClassID:     classMap["Chocolate Meio Amargo"],
				Stock:       50,
				ValTrib:     0.5,
			},
			{
				Code:        "CHBR01",
				Price:       13.00,
				Name:        "Chocolate Branco 90g",
				NCM:         "17049020",
				GTIN:        "7891000000028",
				UM:          "UN",
				Description: "Barra de chocolate branco 90g",
				ClassID:     classMap["Chocolate Branco"],
				Stock:       40,
				ValTrib:     0.5,
			},
			{
				Code:        "BOL01",
				Price:       5.00,
				Name:        "Bolacha Recheada 120g",
				NCM:         "19053100",
				GTIN:        "7891000000035",
				UM:          "UN",
				Description: "Bolacha recheada sabor chocolate 120g",
				ClassID:     classMap["Bolacha"],
				Stock:       100,
				ValTrib:     0.2,
			},
			{
				Code:        "BOM01",
				Price:       2.50,
				Name:        "Bombom Sortido",
				NCM:         "17049030",
				GTIN:        "7891000000042",
				UM:          "UN",
				Description: "Bombom sortido unidade",
				ClassID:     classMap["Bombom"],
				Stock:       200,
				ValTrib:     0.1,
			},
			{
				Code:        "FLR01",
				Price:       15.00,
				Name:        "Buquê de Flores",
				NCM:         "06039000",
				GTIN:        "7891000000059",
				UM:          "UN",
				Description: "Buquê de flores naturais",
				ClassID:     classMap["Flores"],
				Stock:       10,
				ValTrib:     1.0,
			},
		}
		if err := db.Create(&products).Error; err != nil {
			return err
		}
	}

	return nil
}
