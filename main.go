package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type User struct {
	Name      string    `json:"name"`
	IsActive  bool      `json:"isActive"`
	Timestamp time.Time `json:"timeStamp"`
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/home", helloWorld).Methods("GET")

	http.ListenAndServe(":8000", router)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name:      "Dummy Username from backend",
		IsActive:  true,
		Timestamp: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(user)
}
