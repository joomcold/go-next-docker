package controllers

import (
	random "github.com/brianvoe/gofakeit/v6"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joomcold/go-next-docker/internal/app/models"
	"github.com/joomcold/go-next-docker/internal/initializers"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	type FormData struct {
		Email                string `json:"email"`
		Password             string `json:"password"`
		PasswordConfirmation string `json:"passwordConfirmation"`
	}

	var (
		db   = initializers.DB
		form FormData
	)

	// Parse body into form
	err := c.BodyParser(&form)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error", "message": "Invalid input", "data": err,
		})
	}

	// Check password matching
	if form.Password != form.PasswordConfirmation {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error", "message": "Password and Password Confirmation do not match",
		})
	}

	// Encrypt password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), 12)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error", "message": "Failed to hash password",
		})
	}

	// Create user
	user := models.User{
		ID:       uuid.New(),
		Name:     random.Name(),
		Email:    form.Email,
		Password: string(hashedPassword),
		Address:  random.Address().Address,
	}

	err = db.Create(&user).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error", "message": "Failed to register", "data": err.Error(),
		})
	}

	// Return user value
	newUser := models.NewUser{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "successful", "message": "Registration successfully", "data": newUser,
	})
}

func Login(c *fiber.Ctx) error {
	type FormData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var (
		// db   = initializers.DB
		form FormData
	)

	// Parse body into form
	err := c.BodyParser(&form)
	if err != nil || form.Email == "" || form.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error", "message": "Invalid input", "data": err,
		})
	}

	return nil
}
