package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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

func InitDatabase() {
	var err error
	db, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}