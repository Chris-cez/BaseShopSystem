package models

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Name        string       `json:"name"`
	Description string       `json:"description"`
	NCM         string       `json:"ncm"`
	Tributacoes []Tributacao `json:"tributacoes" gorm:"foreignKey:ClassID"`
}

func MigrateClass(db *gorm.DB) error {
	return db.AutoMigrate(&Class{})
}
