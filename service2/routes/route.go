package routes

import (
	"app/controller"

	"github.com/gin-gonic/gin"
)

func ServiceTwoRoute(r *gin.Engine) {
	r.GET("/service2", controller.ServiceTwo)
}
