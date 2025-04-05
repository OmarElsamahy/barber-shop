package handlers

import (
	"barber-shop/config"
	"barber-shop/models"
	"github.com/gofiber/fiber/v2"
)

func GetBarbers(c *fiber.Ctx) error {
	var barbers []models.Barber
	config.DB.Find(&barbers)
	return c.JSON(barbers)
}
