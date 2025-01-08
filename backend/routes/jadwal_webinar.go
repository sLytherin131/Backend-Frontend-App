package routes

import (
	"tugaspemjabackendfrontend/backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterJadwalWebinarRoutes(app *fiber.App) {
	webinar := app.Group("/webinar")
	webinar.Get("/", controllers.GetAllWebinar)
	webinar.Get("/:id", controllers.GetWebinarByPanitiaID)
	webinar.Post("/", controllers.CreateWebinar)
	webinar.Put("/:id", controllers.UpdateWebinar)
	webinar.Delete("/:id", controllers.DeleteWebinar)
}
