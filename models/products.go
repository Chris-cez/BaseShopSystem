package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Code        *string        `json:"code"`
	Price       *float32       `json:"price"`
	Name        *string        `json:"name"`
	NCM         *string        `json:"ncm"`
	UM          *string        `json:"um"`
	Description *string        `json:"description"`
	Class       *uint          `json:"id_class"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func MigrateProducts(db *gorm.DB) error {
	err := db.AutoMigrate(&Product{})
	return err
}
