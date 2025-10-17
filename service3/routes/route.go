package routes

import (
	"app/controller"

	"github.com/gin-gonic/gin"
)

func ServiceThreeRoute(r *gin.Engine) {
	r.GET("/service3", controller.ServiceThree)
}
