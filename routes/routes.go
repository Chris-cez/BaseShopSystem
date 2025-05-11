package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	productRepo := ProductRepository{DB: db}
	companyRepo := CompanyRepository{DB: db}
	addressRepo := AddressRepository{DB: db}
	classRepo := ClassRepository{DB: db}
	clientRepo := ClientRepository{DB: db}
	invoiceItemRepo := InvoiceItemRepository{DB: db}
	invoiceRepo := InvoiceRepository{DB: db}
	paymentMethodRepo := PaymentMethodRepository{DB: db}
	tributacaoRepo := TributacaoRepository{DB: db}

	productRepo.SetupProductRoutes(app)
	addressRepo.SetupAddressRoutes(app)
	classRepo.SetupClassRoutes(app)
	companyRepo.SetupCompanyRoutes(app)
	clientRepo.SetupClientRoutes(app)
	invoiceItemRepo.SetupInvoiceItemRoutes(app)
	invoiceRepo.SetupInvoiceRoutes(app)
	paymentMethodRepo.SetupPaymentMethodRoutes(app)
	tributacaoRepo.SetupTributacaoRoutes(app)
}
