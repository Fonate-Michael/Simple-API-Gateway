package main

import (
	"app/proxy"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Reverse Proxy started and ready to foward the requests hehehe!")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env")

	}

	Service_One := os.Getenv("SERVICE_ONE")
	Service_Two := os.Getenv("SERVICE_TWO")
	Service_Three := os.Getenv("SERVICE_THREE")

	router := gin.Default()
	router.Any("/api1/*path", proxy.ReverseProxy(Service_One, "/api1"))
	router.Any("/api2/*path", proxy.ReverseProxy(Service_Two, "/api2"))
	router.Any("/api3/*path", proxy.ReverseProxy(Service_Three, "/api3"))
	log.Println("API Gate way is running on port 8000 hehehe...")
	router.Run(":8000")

}
