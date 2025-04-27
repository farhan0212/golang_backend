package routes

import (
	"go-crud/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	route := app.Group("/users")
	route.Get("/", controllers.GetUsers)
	route.Post("/", controllers.CreateUser)
	route.Put("/:id", controllers.UpdateUser)
	route.Delete("/:id", controllers.DeleteUser)
}