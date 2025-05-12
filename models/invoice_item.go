package models

import (
	"gorm.io/gorm"
)

type InvoiceItem struct {
	gorm.Model
	ProductID  uint    `json:"product_id"`
	InvoiceID  uint    `json:"invoice_id"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
	ValorTotal float64 `json:"valor_total"`
}

func MigrateInvoiceItem(db *gorm.DB) error {
	return db.AutoMigrate(&InvoiceItem{})
}
