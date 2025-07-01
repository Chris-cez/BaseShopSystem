package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model  `swaggerignore:"true"`
	Logradouro  string `json:"logradouro"`
	Numero      string `json:"numero"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Municipio   string `json:"municipio"`
	UF          string `json:"uf"`
	CEP         string `json:"cep"`
}

func MigrateAddress(db *gorm.DB) error {
	return db.AutoMigrate(&Address{})
}
