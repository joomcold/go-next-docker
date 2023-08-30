package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EncryptPassword(pw string, pwc string, c *gin.Context) {
	if pw != pwc {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password and Password Confirmation do not match"})

		return
	}
}
