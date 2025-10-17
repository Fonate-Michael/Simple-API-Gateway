package routes

import (
	"app/controller"

	"github.com/gin-gonic/gin"
)

func ServiceOneRoute(r *gin.Engine) {
	r.GET("/service1", controller.ServiceOne)
}
