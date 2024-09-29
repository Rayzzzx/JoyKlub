package main

import (
	"log"
	"net/http"

	"github.com/Rayzzzx/JoyKlub/internal/handlers"
	"github.com/Rayzzzx/JoyKlub/internal/repository"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	repository.ConnectDB()

	r := mux.NewRouter()

	// User routes
	r.HandleFunc("/api/register", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/api/login", handlers.LoginUser).Methods("POST")
	r.HandleFunc("/api/check-user", handlers.CheckUserExists).Methods("GET")
	r.HandleFunc("/api/users", handlers.GetAllUsers).Methods("GET") 
	// TODO: Add more routes

	// CORS middleware
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:3000"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    })

    handler := c.Handler(r)

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
