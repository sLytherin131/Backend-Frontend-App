package routes

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	RegisterPanitiaRoutes(app)
	RegisterTugasPanitiaRoutes(app)
	RegisterJadwalWebinarRoutes(app)
}
