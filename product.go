package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code        string  `json:"code"`
	Price       float32 `json:"price"`
	Name        string  `json:"name"`
	NCM         string  `json:"ncm"`
	UM          string  `json:"um"`
	Description string  `json:"description"`
	ClassID     uint    `json:"class_id"`
	Class       Class   `gorm:"foreignKey:ClassID"`
	Stock       int     `json:"stock"`
}

func MigrateProduct(db *gorm.DB) error {
	return db.AutoMigrate(&Product{})
}
