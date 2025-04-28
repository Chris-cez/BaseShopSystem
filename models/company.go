package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	name              string `json:"name" gorm:"not null"`
	CNPJ              string `json:"cnpj" gorm:"unique;not null"`
	inscricaoEstadual string `json:"inscricao_estadual" gorm:"not null"`
	senha             string `json:"senha" gorm:"not null"`
}

func MigrateCompany(db *gorm.DB) error {
	return db.AutoMigrate(&Company{})
}
