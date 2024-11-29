package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name      string  `json:"name"`
	CPF       string  `json:"cpf"`
	AddressID uint    `json:"address_id"`
	Address   Address `gorm:"foreignKey:AddressID"`
}

func MigrateClient(db *gorm.DB) error {
	return db.AutoMigrate(&Client{})
}
