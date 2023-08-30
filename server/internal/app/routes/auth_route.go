package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joomcold/go-next-docker/internal/app/controllers"
)

func Auth(r *gin.Engine) {
	r.POST("/register", controllers.Register)
}
