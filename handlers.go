package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type user struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Points   int    `json:"points"`
}

type allUsers []user

var users = allUsers{
	{
		ID:       "1",
		Username: "Learnable_User",
		Points:   1000,
	},
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Learnable API!")
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser user
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter data for your new user")
	}

	json.Unmarshal(reqBody, &newUser)

	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)
}

func oneUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	for _, oneUser := range users {
		if oneUser.ID == userID {
			json.NewEncoder(w).Encode(oneUser)
		}
	}
}

func allUsersHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}
