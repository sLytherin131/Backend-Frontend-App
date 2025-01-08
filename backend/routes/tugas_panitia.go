package routes

import (
	"github.com/gofiber/fiber/v2"
	"tugaspemjabackendfrontend/backend/controllers"
)

func RegisterTugasPanitiaRoutes(app *fiber.App) {
	tugas := app.Group("/tugas")
	tugas.Get("/", controllers.GetAllTugas)
	tugas.Get("/:id", controllers.GetTugasByPanitiaID)
	tugas.Post("/", controllers.CreateTugas)
	tugas.Put("/:id", controllers.UpdateTugas)
	tugas.Delete("/:id", controllers.DeleteTugas)
}
