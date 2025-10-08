package route

import (
	"backendgo/app/service"
	"backendgo/middleware"

	"github.com/gofiber/fiber/v2"
)

func PekerjaanRoute(api fiber.Router) {
	pekerjaan := api.Group("/pekerjaan")

	pekerjaan.Get("/", middleware.AuthRequired(), service.GetAllPekerjaanService)
	pekerjaan.Get("/list", middleware.AuthRequired(), service.GetAllPekerjaanPaginationService)
	pekerjaan.Get("/trashed", middleware.AuthRequired(), service.GetTrashedPekerjaanService)
	pekerjaan.Put("/:id/restore", middleware.AuthRequired(), middleware.AdminOnly(), service.RestorePekerjaanService)

	pekerjaan.Get("/alumni/:alumni_id", middleware.AuthRequired(), middleware.AdminOnly(), service.GetPekerjaanByAlumniIDService)
	
	pekerjaan.Get("/:id", middleware.AuthRequired(), service.GetPekerjaanByIDService)
	pekerjaan.Post("/", middleware.AuthRequired(), middleware.AdminOnly(), service.CreatePekerjaanService)
	pekerjaan.Put("/:id", middleware.AuthRequired(), middleware.AdminOnly(), service.UpdatePekerjaanService)
	pekerjaan.Delete("/:id", middleware.AuthRequired(), middleware.AdminOnly(), service.DeletePekerjaanService)
	pekerjaan.Put("/:id/soft-delete", middleware.AuthRequired(), service.SoftDeletePekerjaanService)
	pekerjaan.Delete("/:id/hard-delete", middleware.AuthRequired(), middleware.AdminOnly(), service.HardDeletePekerjaanService)

	
}
