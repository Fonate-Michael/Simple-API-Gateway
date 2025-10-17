package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServiceThree(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Service Three"})
}
