package route

import (
	"backendgo/app/service"
	"backendgo/middleware"

	"github.com/gofiber/fiber/v2"
)

func PekerjaanRoute(api fiber.Router) {
	pekerjaan := api.Group("/pekerjaan")

	// 🔹 READ (statis & spesifik dulu)
	pekerjaan.Get("/", middleware.AuthRequired(), service.GetAllPekerjaanService)
	pekerjaan.Get("/list", middleware.AuthRequired(), service.GetAllPekerjaanPaginationService)
	pekerjaan.Get("/trashed", middleware.AuthRequired(), service.GetTrashedPekerjaanService)
	pekerjaan.Get("/alumni/:alumni_id", middleware.AuthRequired(), middleware.AdminOnly(), service.GetPekerjaanByAlumniIDService)
	pekerjaan.Get("/:id", middleware.AuthRequired(), service.GetPekerjaanByIDService)

	// 🔹 CREATE & UPDATE (khusus admin)
	pekerjaan.Post("/", middleware.AuthRequired(), middleware.AdminOnly(), service.CreatePekerjaanService)
	pekerjaan.Put("/:id", middleware.AuthRequired(), middleware.AdminOnly(), service.UpdatePekerjaanService)
	pekerjaan.Delete("/:id", middleware.AuthRequired(), middleware.AdminOnly(), service.DeletePekerjaanService)

	// 🔹 TRASH, RESTORE, HARD DELETE (pakai JWT role di service)
	pekerjaan.Put("/:id/soft-delete", middleware.AuthRequired(), service.SoftDeletePekerjaanService)
	pekerjaan.Put("/:id/restore", middleware.AuthRequired(), service.RestorePekerjaanService)
	pekerjaan.Delete("/:id/hard-delete", middleware.AuthRequired(), service.HardDeletePekerjaanService)
}
