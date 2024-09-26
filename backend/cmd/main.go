package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// TODO: Add routes and handlers

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}