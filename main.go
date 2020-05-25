package main

import (
	"fmt"
	"log"
	"os"

	routes "learnable-be/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := getPort()

	router := gin.Default()

	routes.GetRoutes(router)

	log.Fatal(router.Run(port))
}

func getPort() string {
	var port = os.Getenv("PORT")

	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}

	return ":" + port
}
