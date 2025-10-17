package main

import (
	"app/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Service3 Started and running on port 8003")
	router := gin.Default()
	routes.ServiceThreeRoute(router)
	router.Run(":8003")

}
