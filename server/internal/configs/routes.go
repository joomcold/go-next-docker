package configs

import (
	"github.com/gin-gonic/gin"
	"github.com/joomcold/go-next-docker/internal/app/routes"
)

func Routes() {
	router := gin.Default()

	// Assign routes
	routes.Auth(router)

	err := router.Run()
	if err != nil {
		panic(err)
	}
}
