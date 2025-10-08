package service

import (
	"backendgo/app/model"
	"backendgo/app/repository"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GET semua pekerjaan
func GetAllPekerjaanService(c *fiber.Ctx) error {
	data, err := repository.GetAllPekerjaan()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"success": false, "error": err.Error()})
	}
	return c.JSON(fiber.Map{"success": true, "data": data})
}

// GET pekerjaan by ID
func GetPekerjaanByIDService(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	data, err := repository.GetPekerjaanByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Pekerjaan tidak ditemukan"})
	}
	return c.JSON(fiber.Map{"success": true, "data": data})
}

// GET pekerjaan by Alumni ID
func GetPekerjaanByAlumniIDService(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("alumni_id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid alumni_id"})
	}
	data, err := repository.GetPekerjaanByAlumniID(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"success": true, "data": data})
}

// CREATE pekerjaan
func CreatePekerjaanService(c *fiber.Ctx) error {
	var req model.CreatePekerjaanRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Body tidak valid", "detail": err.Error()})
	}

	mulai, err := time.Parse("2006-01-02", req.TanggalMulaiKerja)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Format tanggal_mulai_kerja harus YYYY-MM-DD"})
	}

	var selesai *time.Time
	if req.TanggalSelesaiKerja != nil && *req.TanggalSelesaiKerja != "" {
		t, err := time.Parse("2006-01-02", *req.TanggalSelesaiKerja)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Format tanggal_selesai_kerja harus YYYY-MM-DD"})
		}
		selesai = &t
	}

	data := model.PekerjaanAlumni{
		AlumniID:            req.AlumniID,
		NamaPerusahaan:      req.NamaPerusahaan,
		PosisiJabatan:       req.PosisiJabatan,
		BidangIndustri:      req.BidangIndustri,
		LokasiKerja:         req.LokasiKerja,
		GajiRange:           req.GajiRange,
		TanggalMulaiKerja:   mulai,
		TanggalSelesaiKerja: selesai,
		StatusPekerjaan:     req.StatusPekerjaan,
		DeskripsiPekerjaan:  req.DeskripsiPekerjaan,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}

	newData, err := repository.CreatePekerjaan(data)
	if err != nil {
		log.Println("Service error CreatePekerjaan:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Gagal insert", "detail": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"success": true, "data": newData})
}

// UPDATE pekerjaan
func UpdatePekerjaanService(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid id"})
	}

	var req model.UpdatePekerjaanRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Body tidak valid"})
	}

	mulai, err := time.Parse("2006-01-02", req.TanggalMulaiKerja)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Format tanggal_mulai_kerja harus YYYY-MM-DD"})
	}

	var selesai *time.Time
	if req.TanggalSelesaiKerja != nil && *req.TanggalSelesaiKerja != "" {
		t, err := time.Parse("2006-01-02", *req.TanggalSelesaiKerja)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Format tanggal_selesai_kerja harus YYYY-MM-DD"})
		}
		selesai = &t
	}

	data := model.PekerjaanAlumni{
		ID:                  id,
		NamaPerusahaan:      req.NamaPerusahaan,
		PosisiJabatan:       req.PosisiJabatan,
		BidangIndustri:      req.BidangIndustri,
		LokasiKerja:         req.LokasiKerja,
		GajiRange:           req.GajiRange,
		TanggalMulaiKerja:   mulai,
		TanggalSelesaiKerja: selesai,
		StatusPekerjaan:     req.StatusPekerjaan,
		DeskripsiPekerjaan:  req.DeskripsiPekerjaan,
		UpdatedAt:           time.Now(),
	}

	updated, err := repository.UpdatePekerjaan(data)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal update"})
	}
	return c.JSON(fiber.Map{"success": true, "data": updated})
}

// DELETE pekerjaan
func DeletePekerjaanService(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := repository.DeletePekerjaan(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"success": true, "message": "Pekerjaan berhasil dihapus"})
}

// GET pagination
func GetAllPekerjaanPaginationService(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	sortBy := c.Query("sortBy", "created_at")
	order := c.Query("order", "desc")
	search := c.Query("search", "")

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * limit

	data, err := repository.GetAllPekerjaanWithPagination(search, sortBy, order, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	total, err := repository.CountPekerjaan(search)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"success":    true,
		"data":       data,
		"page":       page,
		"limit":      limit,
		"total_data": total,
		"total_page": (total + limit - 1) / limit,
	})
}

func SoftDeletePekerjaanService(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid id"})
	}

	role := c.Locals("role").(string)
	userID := c.Locals("user_id").(int)

	if role == "alumni" {
		allowed, err := repository.CheckPekerjaanOwnedByAlumni(id, userID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		if !allowed {
			return c.Status(403).JSON(fiber.Map{"error": "Tidak boleh hapus pekerjaan milik alumni lain"})
		}
	}

	err = repository.SoftDeletePekerjaan(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("Soft delete pekerjaan id %d sukses", id),
	})
}

func GetTrashedPekerjaanService(c *fiber.Ctx) error {
	data, err := repository.GetTrashedPekerjaan()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func RestorePekerjaanService(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "ID tidak valid",
		})
	}

	err = repository.RestorePekerjaan(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Data pekerjaan berhasil direstore",
	})
}

func HardDeletePekerjaanService(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "ID tidak valid",
		})
	}

	err = repository.HardDeletePekerjaan(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Data berhasil dihapus permanen",
	})
}
