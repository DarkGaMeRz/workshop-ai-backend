package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

var db *gorm.DB

func initDatabase() {
	var err error
	db, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate the schema
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}

// GET /users - Get all users
func getUsers(c *fiber.Ctx) error {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}
	return c.JSON(users)
}

// GET /users/:id - Get user by ID
func getUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	
	result := db.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch user",
		})
	}
	
	return c.JSON(user)
}

// POST /users - Create new user
func createUser(c *fiber.Ctx) error {
	var user User
	
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	
	// Validate required fields
	if user.FirstName == "" || user.LastName == "" || user.PhoneNumber == "" || user.Email == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "First name, last name, phone number, and email are required",
		})
	}
	
	// Set registration date to current time
	user.RegistrationDate = time.Now()
	
	result := db.Create(&user)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}
	
	return c.Status(201).JSON(user)
}

// PUT /users/:id - Update user
func updateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	
	// Check if user exists
	result := db.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch user",
		})
	}
	
	// Parse request body
	var updateData User
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	
	// Update user fields
	if updateData.FirstName != "" {
		user.FirstName = updateData.FirstName
	}
	if updateData.LastName != "" {
		user.LastName = updateData.LastName
	}
	if updateData.PhoneNumber != "" {
		user.PhoneNumber = updateData.PhoneNumber
	}
	if updateData.Email != "" {
		user.Email = updateData.Email
	}
	if updateData.MembershipLevel != "" {
		user.MembershipLevel = updateData.MembershipLevel
	}
	if updateData.PointsBalance >= 0 {
		user.PointsBalance = updateData.PointsBalance
	}
	
	result = db.Save(&user)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}
	
	return c.JSON(user)
}

// DELETE /users/:id - Delete user
func deleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	
	// Check if user exists
	result := db.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch user",
		})
	}
	
	// Delete user
	result = db.Delete(&user)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete user",
		})
	}
	
	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}

func main() {
	// Initialize database
	initDatabase()
	
	app := fiber.New()
	
	// Add CORS middleware
	app.Use(cors.New())

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Users API Server is running!")
	})
	
	// User CRUD routes
	app.Get("/users", getUsers)
	app.Get("/users/:id", getUserByID)
	app.Post("/users", createUser)
	app.Put("/users/:id", updateUser)
	app.Delete("/users/:id", deleteUser)

	log.Println("Server starting on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
