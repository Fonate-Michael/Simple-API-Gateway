package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServiceOne(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Service One"})
}
