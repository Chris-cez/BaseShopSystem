package models

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	NCM         string `json:"ncm"`
}

func MigrateClass(db *gorm.DB) error {
	return db.AutoMigrate(&Class{})
}
