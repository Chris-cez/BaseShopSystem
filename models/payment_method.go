package models

import (
	"time"

	"gorm.io/gorm"
)

type PaymentMethod struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func MigratePaymentMethod(db *gorm.DB) error {
	return db.AutoMigrate(&PaymentMethod{})
}
