package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code        string  `json:"code"`
	Price       float32 `json:"price"`
	Name        string  `json:"name"`
	GTIN        string  `json:"gtin"`
	UM          string  `json:"um"`
	Description string  `json:"description"`
	ClassID     uint    `json:"class_id"`
	Stock       int     `json:"stock"`
	ValTrib     float32 `json:"vtribute"`
}

func MigrateProduct(db *gorm.DB) error {
	return db.AutoMigrate(&Product{})
}
