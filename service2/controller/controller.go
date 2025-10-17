package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServiceTwo(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Service Two"})
}
