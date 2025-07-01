package models

import (
	"gorm.io/gorm"
)

type InvoiceItem struct {
	gorm.Model `swaggerignore:"true"`
	ProductID  uint    `json:"product_id"`
	InvoiceID  string  `json:"invoice_id"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
	ValorTotal float64 `json:"valor_total"`
}

func MigrateInvoiceItem(db *gorm.DB) error {
	return db.AutoMigrate(&InvoiceItem{})
}
