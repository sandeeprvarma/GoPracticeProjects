package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"github.com/sandeeprvarma/go-fiber-postgres/models"
	"github.com/sandeeprvarma/go-fiber-postgres/storage"
)

type DBInstance struct {
	DB *gorm.DB	
}

func main()  {
	err	:= godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := storage.NewConnection()
	models.MigrateBooks(db)
	r := DBInstance{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}

func (r *DBInstance)SetupRoutes(app *fiber.App) {
	api :=app.Group("/api")
	api.Get("/books", r.GetBooks)
	// api.Get("/book/:id", r.GetBook)
	// api.Post("/book", r.NewBook)
	// api.Delete("/book/:id", r.DeleteBook)
	// api.Put("/book/:id", r.UpdateBook)
}

func (r *DBInstance)GetBooks(c *fiber.Ctx) error {
	var books []models.Book
	result := r.DB.Find(&books)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(books)
}