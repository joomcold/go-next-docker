package controllers

import (
	"net/http"

	random "github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joomcold/go-next-docker/internal/app/controllers/helpers"
	"github.com/joomcold/go-next-docker/internal/app/models"
	"github.com/joomcold/go-next-docker/internal/initializers"
)

func Register(c *gin.Context) {
	type FormData struct {
		Email                string `form:"email"`
		Password             string `form:"password"`
		PasswordConfirmation string `form:"password_confirmation"`
	}

	var (
		db   = initializers.DB
		form FormData
	)

	err := c.Bind(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})

		return
	}

	helpers.EncryptPassword(form.Password, form.PasswordConfirmation, c)

	user := models.User{
		ID:                   uuid.New(),
		Name:                 random.Name(),
		Email:                form.Email,
		Password:             form.Password,
		PasswordConfirmation: form.PasswordConfirmation,
		Address:              random.Address().Address,
	}

	object := db.Create(&user)
	if object.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
