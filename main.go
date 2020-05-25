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
	// gin.SetMode(gin.ReleaseMode)

	database.Connect()

	router := gin.Default()

	routes.GetRoutes(router)

	log.Fatal(router.Run(port))
}

func getPort() string {
	// Checks the local environment for an environment variable called "PORT"
	var port = os.Getenv("PORT")

	// If the port doesn't exist, it will assign :4747 as the port number
	//    If it does, like in Heroku, then it will use whatever is defined there
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}

	// Returns the port number appended with a colon(:).
	return ":" + port
}
