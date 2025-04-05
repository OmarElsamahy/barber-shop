package main

import (
	"barber-shop/config"
	"barber-shop/models"
	"barber-shop/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.ConnectDB()

	config.DB.AutoMigrate(&models.Barber{})

	routes.BarberRoutes(app)

	app.Listen(":3000")
}
