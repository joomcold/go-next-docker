package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joomcold/go-next-docker/internal/app/models"
	"github.com/joomcold/go-next-docker/internal/initializers"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	type FormData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var (
		db   = initializers.DB
		form FormData
		user models.User
	)

	// Parse body into form
	err := c.BodyParser(&form)
	if err != nil || form.Email == "" || form.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error", "message": "Invalid input",
		})
	}

	// Find user by email
	db.Find(&user, "email = ?", form.Email)

	// Check user email
	if user.Email == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "error", "message": "Incorrect Email or Password",
		})
	}

	// Check user password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error", "message": "Incorrect Email or Password",
		})
	}

	// Create JWT token
	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    user.ID.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	})

	token, err := claim.SignedString([]byte("mango"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error", "message": "Could not login",
		})
	}

	// Attach in cookie
	cookie := fiber.Cookie{
		Name:     "jwt_token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "successful", "message": "Login successfully",
	})
}
