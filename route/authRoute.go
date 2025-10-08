package route

import (
	"backendgo/app/service"
	"backendgo/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(api fiber.Router) {
	api.Post("/login", service.LoginService)

	protected := api.Group("", middleware.AuthRequired())
	protected.Get("/profile", service.GetProfileService) 
}
