package storage

import (
	"github.com/Chris-cez/BaseShopSystem/models"
	"gorm.io/gorm"
	"log"
)

func MigrateModels(db *gorm.DB) error {
	err := models.MigrateProduct(db)
	if err != nil {
		log.Fatal("Error migrating product model\n", err)
	}

	err = models.MigrateClient(db)
	if err != nil {
		log.Fatal("Error migrating client model\n", err)
	}

	err = models.MigrateAddress(db)
	if err != nil {
		log.Fatal("Error migrating address model\n", err)
	}

	err = models.MigrateCompany(db)
	if err != nil {
		log.Fatal("Error migrating company model\n", err)
	}

	err = models.MigrateClass(db)
	if err != nil {
		log.Fatal("Error migrating class model\n", err)
	}

	err = models.MigratePaymentMethod(db)
	if err != nil {
		log.Fatal("Error migrating payment method model\n", err)
	}

	err = models.MigrateInvoice(db)
	if err != nil {
		log.Fatal("Error migrating invoice model\n", err)
	}

	err = models.MigrateInvoiceItem(db)
	if err != nil {
		log.Fatal("Error migrating invoice item model\n", err)
	}

	return nil
}
