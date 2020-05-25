package main

import (
	"fmt"
	"log"
	"os"

	database "learnable-be/config"
	routes "learnable-be/routes"
	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

func main() {
	// Dynamically dictates port for use with Heroku and other services
	port := getPort()

	// Connects to the postgres database
	database.Connect()

	// Start the gin Default router.
	//    Starts a gin.Engine() instance with Logging and Recovery function
	router := gin.Default()

	// Retrieves all routes based on the routes local package
	routes.GetRoutes(router)

	// Logs a fatal response or will run the gin Default router with the port
	//   created by the local getPort function
	//   This is the normal gin way of error handling this process
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
