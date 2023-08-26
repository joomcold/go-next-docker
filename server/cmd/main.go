package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joomcold/go-next-chat/internal/config"
	"github.com/joomcold/go-next-chat/internal/database"
)

func init() {
	config.Environment()

	DB, err := config.Postgresql()
	if err != nil {
		panic(err)
	}

	database.Migrations(DB)
}

func main() {
	route := gin.Default()

	route.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err := route.Run()
	if err != nil {
		panic(err)
	}
}
