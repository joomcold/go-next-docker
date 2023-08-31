package controllers

import (
	"net/http"

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

	err := c.BodyParser(&form)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": "error", "message": "Invalid input", "data": err,
		})
	}

	if form.Password != form.PasswordConfirmation {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": "error", "message": "Password and Password Confirmation do not match",
		})
	}

	cost := 10

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), cost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "error", "message": "Failed to hash password",
		})
	}

	user := models.User{
		ID:       uuid.New(),
		Name:     random.Name(),
		Email:    form.Email,
		Password: string(hashedPassword),
		Address:  random.Address().Address,
	}

	err = db.Create(&user).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "error", "message": "Failed to register", "data": err.Error(),
		})
	}

	newUser := models.NewUser{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "successful", "message": "Registration successfully", "data": newUser,
	})
}
