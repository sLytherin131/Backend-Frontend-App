package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
    "tugaspemjabackendfrontend/backend/database"
    "tugaspemjabackendfrontend/backend/routes"
)

func main() {
    // Initialize database connection
    database.Connect()  // Perbaiki nama fungsi di sini

    // Initialize Fiber app
    app := fiber.New()

    // Register routes
    routes.RegisterRoutes(app)

    // Start the server
    log.Fatal(app.Listen(":3000"))
}
