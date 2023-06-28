package handlers

import (
	"main/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func LoginHandler(c *fiber.Ctx, db *gorm.DB) error {
	// Parse the form values
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Perform authentication logic here
	// Query the database to check if the user with the provided email and password exists
	var user models.User
	if err := db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		// User not found or password incorrect
		// Handle the authentication failure (e.g., redirect to login page with error message)
		return c.Redirect("/login-register/Log-in.html?error=invalid_credentials")
	}

	// Authentication successful
	// Redirect the user to the desired page (e.g., dashboard)
	return c.Redirect("/")
}
