package controllers

import (
	"github.com/gofiber/fiber/v2"
	"tugaspemjabackendfrontend/backend/database"
	"tugaspemjabackendfrontend/backend/models"
)

func GetAllWebinar(c *fiber.Ctx) error {
	var webinar []models.JadwalWebinar
	database.DB.Find(&webinar)
	return c.JSON(webinar)
}

func GetWebinarByPanitiaID(c *fiber.Ctx) error {
	id := c.Params("id")
	var webinar []models.JadwalWebinar
	database.DB.Where("panitia_id = ?", id).Find(&webinar)
	return c.JSON(webinar)
}

func CreateWebinar(c *fiber.Ctx) error {
	var webinar models.JadwalWebinar
	if err := c.BodyParser(&webinar); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	database.DB.Create(&webinar)
	return c.JSON(webinar)
}

func UpdateWebinar(c *fiber.Ctx) error {
	id := c.Params("id")
	var webinar models.JadwalWebinar
	if err := database.DB.First(&webinar, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Webinar not found"})
	}
	if err := c.BodyParser(&webinar); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	database.DB.Save(&webinar)
	return c.JSON(webinar)
}

func DeleteWebinar(c *fiber.Ctx) error {
	id := c.Params("id")
	var webinar models.JadwalWebinar
	if err := database.DB.First(&webinar, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Webinar not found"})
	}
	database.DB.Delete(&webinar)
	return c.JSON(fiber.Map{"message": "Webinar deleted"})
}
