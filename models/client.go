package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model `swaggerignore:"true"`
	Name       string `json:"name"`
	CPF        string `json:"cpf"`
	AddressID  uint   `json:"address_id"`
}

func MigrateClient(db *gorm.DB) error {
	return db.AutoMigrate(&Client{})
}
