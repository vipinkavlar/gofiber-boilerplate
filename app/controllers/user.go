package controller

import (
	"github.com/gofiber/fiber/v2"
)

func GetUserByID(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"code": 200, "message": "Pong!", "data" : "User details"})
}