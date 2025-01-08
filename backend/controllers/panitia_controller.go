package controllers

import (
	"github.com/gofiber/fiber/v2"
	"tugaspemjabackendfrontend/backend/database"
	"tugaspemjabackendfrontend/backend/models"
	"strconv"
)

func GetAllPanitia(c *fiber.Ctx) error {
	var panitia []models.Panitia
	if err := database.DB.Find(&panitia).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch data"})
	}
	return c.JSON(panitia)
}

func CreatePanitia(c *fiber.Ctx) error {
	var panitia models.Panitia
	if err := c.BodyParser(&panitia); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := database.DB.Create(&panitia).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create panitia"})
	}
	return c.Status(201).JSON(panitia)
}

func UpdatePanitia(c *fiber.Ctx) error {
	id := c.Params("id") // Mengambil id dari parameter URL
	// Validasi ID agar sesuai format
	if _, err := strconv.Atoi(id); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var panitia models.Panitia
	// Menggunakan id_panitia sebagai kolom pencarian
	if err := database.DB.First(&panitia, "id_panitia = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Panitia not found"})
	}
	if err := c.BodyParser(&panitia); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := database.DB.Save(&panitia).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update panitia"})
	}
	return c.JSON(panitia)
}

func DeletePanitia(c *fiber.Ctx) error {
	id := c.Params("id") // Mengambil id dari parameter URL
	// Validasi ID agar sesuai format
	if _, err := strconv.Atoi(id); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var panitia models.Panitia
	// Menggunakan id_panitia sebagai kolom pencarian
	if err := database.DB.First(&panitia, "id_panitia = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Panitia not found"})
	}
	if err := database.DB.Delete(&panitia).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete panitia"})
	}
	return c.JSON(fiber.Map{"message": "Panitia deleted"})
}
