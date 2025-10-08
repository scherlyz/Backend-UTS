package service

import (
	"backendgo/app/model"
	"backendgo/app/repository"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

// GET semua alumni
func GetAllAlumniService(c *fiber.Ctx) error {
	data, err := repository.GetAllAlumni()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"success": true, "data": data})
}

// GET alumni by ID
func GetAlumniByIDService(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
	}
	data, err := repository.GetAlumniByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Alumni tidak ditemukan"})
	}
	return c.JSON(fiber.Map{"success": true, "data": data})
}

// CREATE alumni baru
func CreateAlumniService(c *fiber.Ctx) error {
	var input model.Alumni
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	data, err := repository.CreateAlumni(input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"success": true, "data": data})
}

// UPDATE alumni
func UpdateAlumniService(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
	}
	var input model.Alumni
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	input.ID = id
	input.UpdatedAt = time.Now()

	data, err := repository.UpdateAlumni(input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"success": true, "data": data})
}

// DELETE alumni
func DeleteAlumniService(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
	}
	if err := repository.DeleteAlumni(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"success": true, "message": "Alumni berhasil dihapus"})
}

// UPDATE status kematian
func UpdateStatusKematianService(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
	}

	var req struct {
		StatusKematian bool `json:"status_kematian"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Body tidak valid"})
	}

	if err := repository.UpdateStatusKematian(id, req.StatusKematian); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal update status kematian"})
	}
	return c.JSON(fiber.Map{"success": true, "message": "Status kematian diperbarui"})
}

// GET alumni with pagination
func GetAlumniWithPaginationService(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	sortBy := c.Query("sortBy", "id")
	order := c.Query("order", "asc")
	search := c.Query("search", "")

	offset := (page - 1) * limit

	data, err := repository.GetAlumniRepo(search, sortBy, order, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	total, _ := repository.CountAlumniRepo(search)

	response := model.AlumniResponse{
		Data: data,
		Meta: model.MetaInfo{
			Page:   page,
			Limit:  limit,
			Total:  total,
			Pages:  (total + limit - 1) / limit,
			SortBy: sortBy,
			Order:  order,
			Search: search,
		},
	}
	return c.JSON(response)
}
