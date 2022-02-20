package engine

import (
	"github.com/gofiber/fiber/v2"
  "github.com/vipinkavlar/gofiber-boilerplate/routes"
)

func Start() {
  app := fiber.New()

  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
  })

  routes.LoadRoutes(app)

  app.Listen(":3000")
}