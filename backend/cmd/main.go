package main

import (
	"log"
	"net/http"

	"ecommerce/backend/internal/repository"

	"github.com/gorilla/mux"
)

func main() {
	repository.ConnectDB()

	r := mux.NewRouter()

	// TODO: Add routes and handlers

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
