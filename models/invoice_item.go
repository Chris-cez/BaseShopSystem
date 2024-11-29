package models

import (
	"gorm.io/gorm"
)

type InvoiceItem struct {
	gorm.Model
	InvoiceID  uint    `json:"invoice_id"`
	ProductID  uint    `json:"product_id"`
	Product    Product `gorm:"foreignKey:ProductID"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
	ValorTotal float64 `json:"valor_total"`
}

func MigrateInvoiceItem(db *gorm.DB) error {
	return db.AutoMigrate(&InvoiceItem{})
}
