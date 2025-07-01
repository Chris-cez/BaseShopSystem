package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model  `swaggerignore:"true"`
	Code        string  `json:"code"`
	Price       float64 `json:"price"`
	Name        string  `json:"name"`
	NCM         string  `json:"ncm"`
	GTIN        string  `json:"gtin"`
	UM          string  `json:"um"`
	Description string  `json:"description"`
	ClassID     uint    `json:"class_id"`
	Stock       int     `json:"stock"`
	ValTrib     float64 `json:"valtrib"`
}

func MigrateProduct(db *gorm.DB) error {
	return db.AutoMigrate(&Product{})
}
