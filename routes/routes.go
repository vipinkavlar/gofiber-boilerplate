package routes

import(
	"github.com/gofiber/fiber/v2"
	"github.com/vipinkavlar/gofiber-boilerplate/app/controllers"
)

func LoadRoutes(a *fiber.App) {
	baseRoute := a.Group("/api")
	v1 := baseRoute.Group("/v1")
    
	//Pinger Route
	v1.Get("/", controller.Ping)

	//User Routes
	userV1 := v1.Group("/user")
	userV1.Get("/getUserById", controller.GetUserByID)
}