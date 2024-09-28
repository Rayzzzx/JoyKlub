package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Rayzzzx/JoyKlub/internal/models"
	"github.com/Rayzzzx/JoyKlub/internal/repository"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Insert user into database
	collection := repository.Client.Database("ecommerce").Collection("users")
	_, err = collection.InsertOne(r.Context(), user)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Find user in database
	collection := repository.Client.Database("ecommerce").Collection("users")
	var user models.User
	err = collection.FindOne(r.Context(), bson.M{"email": loginData.Email}).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// TODO: Generate and return JWT token

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
}
