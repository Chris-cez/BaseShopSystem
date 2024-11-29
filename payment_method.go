package models

import "gorm.io/gorm"

type PaymentMethod struct {
    gorm.Model
    Name string `json:"name"`
}

func MigratePaymentMethod(db *gorm.DB) error {
    return db.AutoMigrate(&PaymentMethod{})
}