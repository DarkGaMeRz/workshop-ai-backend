package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// User model based on the UI fields shown
type User struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	FirstName        string    `json:"first_name" gorm:"not null"`
	LastName         string    `json:"last_name" gorm:"not null"`
	PhoneNumber      string    `json:"phone_number" gorm:"unique;not null"`
	Email            string    `json:"email" gorm:"unique;not null"`
	RegistrationDate time.Time `json:"registration_date" gorm:"autoCreateTime"`
	MembershipLevel  string    `json:"membership_level" gorm:"default:Bronze"`
	PointsBalance    int       `json:"points_balance" gorm:"default:0"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func main() {
	// Initialize database
	InitDatabase()

	app := fiber.New()

	// Middleware
	app.Use(cors.New())
	app.Use(logger.New())

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Users API Server is running!")
	})
	app.Get("/users", GetUsers)
	app.Get("/users/:id", GetUserByID)
	app.Post("/users", CreateUser)
	app.Put("/users/:id", UpdateUser)
	app.Delete("/users/:id", DeleteUser)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("Server starting on port %s...", port)
	log.Fatal(app.Listen(":" + port))
}
