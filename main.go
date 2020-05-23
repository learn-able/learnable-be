package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := GetPort()
	fmt.Println("Starting Service...")

	fmt.Println("Started...")
	fmt.Println("Listening on Port: " + port)

	router := mux.NewRouter()

	router.HandleFunc("/api/v0/user", userHandler).Methods("GET")
	http.Handle("/", router)

	err := http.ListenAndServe(port, router)

	if err != nil {
		panic(err)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {

}

func GetPort() string {
	var port = os.Getenv("PORT")

	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}

	return ":" + port
}
