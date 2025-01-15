package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

// Database connection
var db *sql.DB

func connectDatabase() {
	var err error
	// Ubah konfigurasi koneksi sesuai kebutuhan
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

	// Routes for Fields
	app.Get("/fields", getFields)
	app.Post("/fields", createField)
	app.Put("/fields/:id", updateField)
	app.Delete("/fields/:id", deleteField)

	// Routes for Customers
	app.Get("/customers", getCustomers)
	app.Post("/customers", createCustomer)
	app.Put("/customers/:id", updateCustomer)
	app.Delete("/customers/:id", deleteCustomer)

	// Routes for Reservations
	app.Get("/reservations", getReservations)
	app.Post("/reservations", createReservation)
	app.Put("/reservations/:id", updateReservation)
	app.Delete("/reservations/:id", deleteReservation)

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

	_, err := db.Exec("INSERT INTO fields (name, type, price_per_hour) VALUES (?, ?, ?)",
		field.Name, field.Type, field.PricePerHour)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(201)
}

func updateField(c *fiber.Ctx) error {
	id := c.Params("id")
	var field Field
	if err := c.BodyParser(&field); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := db.Exec("UPDATE fields SET name = ?, type = ?, price_per_hour = ? WHERE id = ?",
		field.Name, field.Type, field.PricePerHour, id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(200)
}

func deleteField(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := db.Exec("DELETE FROM fields WHERE id = ?", id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(200)
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

	_, err := db.Exec("INSERT INTO customers (name, phone_number) VALUES (?, ?)",
		customer.Name, customer.PhoneNumber)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(201)
}

func updateCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	var customer Customer
	if err := c.BodyParser(&customer); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := db.Exec("UPDATE customers SET name = ?, phone_number = ? WHERE id = ?",
		customer.Name, customer.PhoneNumber, id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(200)
}

func deleteCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := db.Exec("DELETE FROM customers WHERE id = ?", id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(200)
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
		var startTimeRaw []uint8 // Temp variable for raw database value

		if err := rows.Scan(&reservation.ID, &reservation.FieldID, &reservation.CustomerID,
			&startTimeRaw, &reservation.Duration, &reservation.TotalCost); err != nil {
			return c.Status(500).SendString(err.Error())
		}

		// Parse startTimeRaw to time.Time
		parsedTime, err := time.Parse("2006-01-02 15:04:05", string(startTimeRaw))
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
		return c.Status(400).SendString(err.Error())
	}

	// Format waktu sebelum mengirim ke database
	startTimeFormatted := reservation.StartTime.Format("2006-01-02 15:04:05")

	_, err := db.Exec("INSERT INTO reservations (field_id, customer_id, start_time, duration, total_cost) VALUES (?, ?, ?, ?, ?)",
		reservation.FieldID, reservation.CustomerID, startTimeFormatted, reservation.Duration, reservation.TotalCost)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(201)
}

func updateReservation(c *fiber.Ctx) error {
	id := c.Params("id")
	var reservation Reservation
	if err := c.BodyParser(&reservation); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := db.Exec("UPDATE reservations SET field_id = ?, customer_id = ?, start_time = ?, duration = ?, total_cost = ? WHERE id = ?",
		reservation.FieldID, reservation.CustomerID, reservation.StartTime, reservation.Duration, reservation.TotalCost, id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(200)
}

func deleteReservation(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := db.Exec("DELETE FROM reservations WHERE id = ?", id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(200)
}
