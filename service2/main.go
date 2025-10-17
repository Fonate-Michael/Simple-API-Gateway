package main

import (
	"app/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Service2 Started and running on port 8002")
	router := gin.Default()
	routes.ServiceTwoRoute(router)
	router.Run(":8002")

}
