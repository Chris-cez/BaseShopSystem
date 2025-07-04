package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model        `swaggerignore:"true"`
	Name              string `json:"name" gorm:"not null"`
	CNPJ              string `json:"cnpj" gorm:"unique;not null"`
	InscricaoEstadual string `json:"inscricao_estadual" gorm:"not null"`
	Password          string `json:"password" gorm:"not null"`
	Address_id        int    `json:"address" gorm:"not null"`
}

func MigrateCompany(db *gorm.DB) error {
	return db.AutoMigrate(&Company{})
}
