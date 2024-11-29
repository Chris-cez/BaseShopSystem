package models

import "gorm.io/gorm"

type Tributacao struct {
	gorm.Model
	Nome        string  `json:"nome"`
	Aliquota    float64 `json:"aliquota"`
	TipoTributo string  `json:"tipo_tributo"`
}

func MigrateTributacao(db *gorm.DB) error {
	return db.AutoMigrate(&Tributacao{})
}
