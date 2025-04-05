package routes

import (
	"barber-shop/handlers"
	"github.com/gofiber/fiber/v2"
)

func BarberRoutes(app *fiber.App) {
	app.Get("/barbers", handlers.GetBarbers)
}
