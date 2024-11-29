package storage

import (
    "github.com/Chris-cez/BaseShopSystem/models"
    "gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) error {
    err := models.MigrateProduct(db)
    if err != nil {
        return err
    }

    err = models.MigrateClient(db)
    if err != nil {
        return err
    }

    err = models.MigrateAddress(db)
    if err != nil {
        return err
    }

    err = models.MigrateClass(db)
    if err != nil {
        return err
    }

    err = models.MigratePaymentMethod(db)
    if err != nil {
        return err
    }

    err = models.MigrateInvoice(db)
    if err != nil {
        return err
    }

    err = models.MigrateInvoiceItem(db)
    if err != nil {
        return err
    }

    err = models.MigrateTributacao(db)
    if err != nil {
        return err
    }

    return nil
}