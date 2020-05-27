package main

import (
	"fmt"
	"log"
	"os"

	database "learnable-be/config"
	routes "learnable-be/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	port := getPort()
	gin.SetMode(gin.ReleaseMode) // set for production

	// LOCAL CONFIG
	// envErr := godotenv.Load("application.yml")
	// if envErr != nil {
	// 	log.Fatal("Error loading .yml file")
	// }
	// END LOCAL CONFIG

	database.Connect()

	router := routes.GetRoutes()

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
