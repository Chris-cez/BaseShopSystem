package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model			
	Code  string			`json: "code"`
	Price float32			`json: "price"`
	Name  string			`json: "name"`
	NCM   string			`json: "ncm"`
	UM    string			`json: "um"`
	Description string		`json: "description"`
	Id_Class int			`json: "id_class"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateProduct(context &fiber.Ctx) error {
	product := Product{}

	err := context.BodyParser(&product)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Could not create product"},
		)

		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Product has been added"},
	)
	return nil
}

func (r *Repository) GetProducts(context &fiber.Ctx) error {
	productModels := []models.Product{}

	err := r.DB.Find(&productModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not get products"},
		)
	}
}

func(r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/products", r.CreateProduct)
	api.Get("/products", r.GetProducts)
	api.Get("/products/:id", r.GetProduct)
	api.Put("/products/:id", r.UpdateProduct)
	api.Delete("/products/:id", r.DeleteProduct)
}

func main()  {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file", err)
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatalf("Error connecting to database", err)
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
