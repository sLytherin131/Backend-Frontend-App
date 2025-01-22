package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var db *sql.DB

func connectDatabase() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/reservasi_futsal")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Database connection error:", err)
	}
	fmt.Println("Connected to the database!")
}

type Field struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Type         string  `json:"type"`
	PricePerHour float64 `json:"price_per_hour"`
}

type Customer struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type Reservation struct {
	ID         int       `json:"id"`
	FieldID    int       `json:"field_id"`
	CustomerID int       `json:"customer_id"`
	StartTime  time.Time `json:"start_time"`
	Duration   int       `json:"duration"`
	TotalCost  float64   `json:"total_cost"`
}

func main() {
	connectDatabase()
	defer db.Close()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	// Routes
	app.Get("/fields", getFields)
	app.Post("/fields", createField)
	app.Put("/fields/:id", updateField)    // Update Field
	app.Delete("/fields/:id", deleteField) // Delete Field

	app.Get("/customers", getCustomers)
	app.Post("/customers", createCustomer)
	app.Put("/customers/:id", updateCustomer)    // Update Customer
	app.Delete("/customers/:id", deleteCustomer) // Delete Customer

	app.Get("/reservations", getReservations)
	app.Post("/reservations", createReservation)
	app.Delete("/reservations/:id", deleteReservation) // Delete Reservation

	log.Fatal(app.Listen(":3000"))
}

// CRUD Handlers for Fields
func getFields(c *fiber.Ctx) error {
	rows, err := db.Query("SELECT id, name, type, price_per_hour FROM fields")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	var fields []Field
	for rows.Next() {
		var field Field
		if err := rows.Scan(&field.ID, &field.Name, &field.Type, &field.PricePerHour); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		fields = append(fields, field)
	}

	return c.JSON(fields)
}

func createField(c *fiber.Ctx) error {
	var field Field
	if err := c.BodyParser(&field); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if field.Name == "" || field.Type == "" || field.PricePerHour <= 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input data"})
	}

	result, err := db.Exec("INSERT INTO fields (name, type, price_per_hour) VALUES (?, ?, ?)",
		field.Name, field.Type, field.PricePerHour)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	id, _ := result.LastInsertId()
	return c.Status(201).JSON(fiber.Map{
		"message": "Field created successfully",
		"id":      id,
	})
}

func updateField(c *fiber.Ctx) error {
	id := c.Params("id")
	var field Field
	if err := c.BodyParser(&field); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if field.Name == "" || field.Type == "" || field.PricePerHour <= 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input data"})
	}

	_, err := db.Exec("UPDATE fields SET name = ?, type = ?, price_per_hour = ? WHERE id = ?",
		field.Name, field.Type, field.PricePerHour, id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendString("Field updated successfully")
}

func deleteField(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := db.Exec("DELETE FROM fields WHERE id = ?", id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendString("Field deleted successfully")
}

// CRUD Handlers for Customers
func getCustomers(c *fiber.Ctx) error {
	rows, err := db.Query("SELECT id, name, phone_number FROM customers")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var customer Customer
		if err := rows.Scan(&customer.ID, &customer.Name, &customer.PhoneNumber); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		customers = append(customers, customer)
	}

	return c.JSON(customers)
}

func createCustomer(c *fiber.Ctx) error {
	var customer Customer
	if err := c.BodyParser(&customer); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if customer.Name == "" || customer.PhoneNumber == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input data"})
	}

	result, err := db.Exec("INSERT INTO customers (name, phone_number) VALUES (?, ?)",
		customer.Name, customer.PhoneNumber)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	id, _ := result.LastInsertId()
	return c.Status(201).JSON(fiber.Map{
		"message": "Customer created successfully",
		"id":      id,
	})
}

func updateCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	var customer Customer
	if err := c.BodyParser(&customer); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if customer.Name == "" || customer.PhoneNumber == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input data"})
	}

	_, err := db.Exec("UPDATE customers SET name = ?, phone_number = ? WHERE id = ?",
		customer.Name, customer.PhoneNumber, id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendString("Customer updated successfully")
}

func deleteCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := db.Exec("DELETE FROM customers WHERE id = ?", id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendString("Customer deleted successfully")
}

// CRUD Handlers for Reservations
func getReservations(c *fiber.Ctx) error {
	rows, err := db.Query("SELECT id, field_id, customer_id, start_time, duration, total_cost FROM reservations")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	var reservations []Reservation
	for rows.Next() {
		var reservation Reservation
		var startTimeRaw string
		if err := rows.Scan(&reservation.ID, &reservation.FieldID, &reservation.CustomerID,
			&startTimeRaw, &reservation.Duration, &reservation.TotalCost); err != nil {
			return c.Status(500).SendString(err.Error())
		}

		parsedTime, err := time.Parse("2006-01-02 15:04:05", startTimeRaw)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		reservation.StartTime = parsedTime

		reservations = append(reservations, reservation)
	}

	return c.JSON(reservations)
}

func createReservation(c *fiber.Ctx) error {
	var reservation Reservation
	if err := c.BodyParser(&reservation); err != nil {
		return c.Status(400).SendString("Invalid input: " + err.Error())
	}

	// Parse waktu mulai dari input
	startTime, err := time.Parse("2006-01-02T15:04:05", reservation.StartTime.Format("2006-01-02T15:04:05"))
	if err != nil {
		return c.Status(400).SendString("Invalid start time format: " + err.Error())
	}

	// Ambil harga lapangan berdasarkan FieldID
	var pricePerHour float64
	err = db.QueryRow("SELECT price_per_hour FROM fields WHERE id = ?", reservation.FieldID).Scan(&pricePerHour)
	if err != nil {
		return c.Status(400).SendString("Field not found")
	}

	// Hitung total biaya
	reservation.TotalCost = pricePerHour * float64(reservation.Duration)

	// Simpan ke database
	_, err = db.Exec("INSERT INTO reservations (field_id, customer_id, start_time, duration, total_cost) VALUES (?, ?, ?, ?, ?)",
		reservation.FieldID, reservation.CustomerID, startTime.Format("2006-01-02 15:04:05"), reservation.Duration, reservation.TotalCost)
	if err != nil {
		return c.Status(500).SendString("Failed to save reservation: " + err.Error())
	}

	return c.Status(201).JSON(fiber.Map{"message": "Reservation created successfully"})
}

func deleteReservation(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := db.Exec("DELETE FROM reservations WHERE id = ?", id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendString("Reservation deleted successfully")
}
