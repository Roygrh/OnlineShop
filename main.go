package main

import (
	"log"
	//"github.com/gofiber/fiber/v2"
	"OnlineShop/cmd/server"
	//"OnlineShop/internal/configs"
	"OnlineShop/pkg/routes"
)

func main() {

	//logging
	log.Println("Starting up...")

	//load configurations
	/*if err := configs.load(); err != nil {
		log.Fatal(err)
	}*/

	app := server.StartServer()

	routes.SetupRoutes(app)

	app.Listen(":8095")

	log.Println("Server started on port 8099")
}