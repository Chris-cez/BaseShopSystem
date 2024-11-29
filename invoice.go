package models

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	Numero          string        `json:"numero" gorm:"primaryKey"`
	ClientID        uint          `json:"client_id"`
	Client          Client        `gorm:"foreignKey:ClientID"`
	Items           []InvoiceItem `json:"items" gorm:"foreignKey:Invoice_ItemID"`
	TotalValue      float64       `json:"total_value"`
	PaymentMethodID uint          `json:"payment_method_id"`
	PaymentMethod   PaymentMethod `gorm:"foreignKey:PaymentMethodID"`
	Discount        float64       `json:"discount"`
	Observation     string        `json:"observation"`
	AccessKey       string        `json:"access_key"`
	CreatedAt       time.Time     `json:"created_at"`
}

func MigrateInvoice(db *gorm.DB) error {
	return db.AutoMigrate(&Invoice{})
}
