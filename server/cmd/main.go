package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joomcold/go-next-docker/internal/database"
	"github.com/joomcold/go-next-docker/internal/initializers"
)

func init() {
	initializers.Environment()
	initializers.Postgresql()
	database.Migrations()
}

func main() {
	route := gin.Default()

	err := route.Run()
	if err != nil {
		panic(err)
	}
}
