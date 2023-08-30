package controllers

import (
	"net/http"

	random "github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joomcold/go-next-docker/internal/app/models"
	"github.com/joomcold/go-next-docker/internal/initializers"
)

func Register(c *gin.Context) {
	var db = initializers.DB

	var body struct {
		Name                 string
		Email                string
		Password             string
		PasswordConfirmation string
		Address              string
	}

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})

		return
	}

	user := models.User{
		ID:      uuid.New(),
		Name:    random.Name(),
		Email:   body.Email,
		Address: random.Address().Address,
	}

	object := db.Create(&user)
	if object.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
