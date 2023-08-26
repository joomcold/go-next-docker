package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joomcold/go-next-chat/internal/config"
)

func init() {
	config.LoadEnv()

	_, err := config.ConnectPsql()
	if err != nil {
		panic(err)
	}
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
