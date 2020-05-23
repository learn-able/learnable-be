package main

import "github.com/gorilla/mux"

func NewRouter(router *mux.Router) {
	router.Methods("GET").Path("/api/v0").HandlerFunc(indexHandler)
	router.Methods("POST").Path("/api/v0/users").HandlerFunc(createUserHandler)
	router.Methods("GET").Path("/api/v0/users").HandlerFunc(allUsersHandler)
	router.Methods("GET").Path("/api/v0/users/{id}").HandlerFunc(oneUserHandler)
}
