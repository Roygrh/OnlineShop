package handlers

import (
	"github.com/gofiber/fiber/v2"
	internalModels "OnlineShop/internal/models"
	pkgModels "OnlineShop/pkg/models"
)

func GetDocumentsList(c *fiber.Ctx) error{
	
	docs := []internalModels.Document{
		{
			Id: 1,
			Name: "Document 1",
			Data: []byte{0x1, 0x2},
		},
		{
			Id: 2,
			Name: "Document 2",
			Data: []byte{0x3, 0x4},
		},
		{
			Id: 3,
			Name: "Document 3",
			Data: []byte{0x5, 0x6},
		},
	}

	return c.JSON(docs)

}

func DocumentAnalysis(c *fiber.Ctx) error {
	docAnalysis := pkgModels.AnalysisResult{
		Name: "RESULTADO DEL ANÁLISIS",
		Conclusion: "Resultado del análisis de los archivos, documentos y fuentes de datos",
		Data: []byte{0x1, 0x2},
	}

	return c.JSON(docAnalysis)
}

func DocumentsSearch(c *fiber.Ctx) error {
	docs := []internalModels.Document{
		{
			Id: 1,
			Name: "Documento buscado 1",
			Data: []byte{0x1, 0x2},
		},
		{
			Id: 2,
			Name: "Documento buscado 2",
			Data: []byte{0x3, 0x4},
		},
		{
			Id: 3,
			Name: "Documento buscado 3",
			Data: []byte{0x5, 0x6},
		},
		{
			Id: 4,
			Name: "Documento buscado 4",
			Data: []byte{0x7, 0x8},
		},
	}

	return c.JSON(docs)
}

func DocumentRegistration(c *fiber.Ctx) error{
	return c.JSON("Se registro un nuevo documento")
}