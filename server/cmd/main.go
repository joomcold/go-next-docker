package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joomcold/go-next-chat/internal/config"
)

func init() {
	config.LoadEnv()
	config.ConnectPsql()
}

func main() {
	route := gin.Default()

	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	route.Run()
}
