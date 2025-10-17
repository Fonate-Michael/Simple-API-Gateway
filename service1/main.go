package main

import (
	"app/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Service1 Started and running on port 8001")
	router := gin.Default()
	routes.ServiceOneRoute(router)
	router.Run(":8001")

}
