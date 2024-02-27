package routes

import (
	"github.com/gofiber/fiber/v2"
	"OnlineShop/pkg/handlers"
)

func SetupRoutes(app * fiber.App){
	app.Get("/documents-list", handlers.GetDocumentsList)
	app.Get("/documents-search", handlers.DocumentsSearch)
	app.Post("/document-analysis", handlers.DocumentAnalysis)
	app.Post("/document-registration", handlers.DocumentRegistration)
}