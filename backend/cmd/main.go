package main

import (
	"log"
	"net/http"

	"ecommerce/backend/internal/handlers"
	"ecommerce/backend/internal/repository"

	"github.com/gorilla/mux"
)

func main() {
	repository.ConnectDB()

	r := mux.NewRouter()

	// User routes
	r.HandleFunc("/api/register", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/api/login", handlers.LoginUser).Methods("POST")

	// TODO: Add more routes

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
