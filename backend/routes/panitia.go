package routes

import (
	"github.com/gofiber/fiber/v2"
	"tugaspemjabackendfrontend/backend/controllers"
)

func RegisterPanitiaRoutes(app *fiber.App) {
	panitia := app.Group("/panitia")
	panitia.Get("/", controllers.GetAllPanitia)
	panitia.Post("/", controllers.CreatePanitia)
	panitia.Put("/:id", controllers.UpdatePanitia)
	panitia.Delete("/:id", controllers.DeletePanitia)
}
