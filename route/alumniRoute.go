package route

import (
	"backendgo/app/service"
	"backendgo/middleware"

	"github.com/gofiber/fiber/v2"
)

func AlumniRoute(api fiber.Router) {
	alumni := api.Group("/alumni")

	alumni.Get("/", middleware.AuthRequired(), service.GetAllAlumniService)
	alumni.Get("/list", middleware.AuthRequired(), service.GetAlumniWithPaginationService)
	alumni.Get("/:id", middleware.AuthRequired(), service.GetAlumniByIDService)
	alumni.Post("/", middleware.AuthRequired(), middleware.AdminOnly(), service.CreateAlumniService)
	alumni.Put("/:id", middleware.AuthRequired(), middleware.AdminOnly(), service.UpdateAlumniService)
	alumni.Delete("/:id", middleware.AuthRequired(), middleware.AdminOnly(), service.DeleteAlumniService)
	alumni.Put("/:id/kematian", middleware.AuthRequired(), middleware.AdminOnly(), service.UpdateStatusKematianService)
}
