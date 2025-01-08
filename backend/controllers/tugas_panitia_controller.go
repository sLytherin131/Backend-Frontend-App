package controllers

import (
	"strconv"
	"tugaspemjabackendfrontend/backend/database"
	"tugaspemjabackendfrontend/backend/models"

	"github.com/gofiber/fiber/v2"
)

// GetAllTugas retrieves all tugas from the database
func GetAllTugas(c *fiber.Ctx) error {
	var tugas []models.TugasPanitia
	if err := database.DB.Find(&tugas).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch tugas"})
	}
	return c.JSON(tugas)
}

// GetTugasByPanitiaID retrieves all tugas by PanitiaID
func GetTugasByPanitiaID(c *fiber.Ctx) error {
	id := c.Params("id")
	var tugas []models.TugasPanitia
	if err := database.DB.Where("panitia_id = ?", id).Find(&tugas).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch tugas"})
	}
	return c.JSON(tugas)
}

func CreateTugas(c *fiber.Ctx) error {
	var tugas models.TugasPanitia
	if err := c.BodyParser(&tugas); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Validasi apakah PanitiaID ada di database
	var panitia models.Panitia
	if err := database.DB.First(&panitia, tugas.PanitiaID).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid PanitiaID"})
	}

	// Validasi jika Deskripsi kosong
	if tugas.Deskripsi == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Deskripsi cannot be empty"})
	}

	// Simpan tugas
	if err := database.DB.Create(&tugas).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create tugas"})
	}
	return c.Status(201).JSON(tugas)
}

// UpdateTugas updates an existing tugas by ID
func UpdateTugas(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var tugas models.TugasPanitia
	if err := database.DB.First(&tugas, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Tugas not found"})
	}

	if err := c.BodyParser(&tugas); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Update tugas
	if err := database.DB.Save(&tugas).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update tugas"})
	}
	return c.JSON(tugas)
}

// DeleteTugas deletes a tugas by ID
func DeleteTugas(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var tugas models.TugasPanitia
	if err := database.DB.First(&tugas, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Tugas not found"})
	}

	// Delete tugas
	if err := database.DB.Delete(&tugas).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete tugas"})
	}
	return c.JSON(fiber.Map{"message": "Tugas deleted successfully"})
}
