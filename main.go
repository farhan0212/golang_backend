package main

import (
	database "go-crud/config"
	"go-crud/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.InitDB()
	routes.SetupRoutes(app)

	log.Println("Server running on port 3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Failed to start server", err)
	}
}