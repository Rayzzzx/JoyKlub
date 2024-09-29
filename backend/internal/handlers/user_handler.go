package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Rayzzzx/JoyKlub/internal/models"
	"github.com/Rayzzzx/JoyKlub/internal/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if user already exists
	collection := repository.Client.Database("ecommerce").Collection("users")
	var existingUser models.User
	err = collection.FindOne(r.Context(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		http.Error(w, "User with this email already exists", http.StatusConflict)
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
	_, err = collection.InsertOne(r.Context(), user)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
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

func CheckUserExists(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Email parameter is required", http.StatusBadRequest)
		return
	}

	collection := repository.Client.Database("ecommerce").Collection("users")
	var user models.User
	err := collection.FindOne(r.Context(), bson.M{"email": email}).Decode(&user)

	w.Header().Set("Content-Type", "application/json")
	if err == mongo.ErrNoDocuments {
		json.NewEncoder(w).Encode(map[string]bool{"exists": false})
	} else if err != nil {
		http.Error(w, "Error checking user existence", http.StatusInternalServerError)
		return
	} else {
		json.NewEncoder(w).Encode(map[string]bool{"exists": true})
	}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	collection := repository.Client.Database("ecommerce").Collection("users")
	
	cursor, err := collection.Find(r.Context(), bson.M{})
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(r.Context())

	var users []models.User
	if err = cursor.All(r.Context(), &users); err != nil {
		http.Error(w, "Error decoding users", http.StatusInternalServerError)
		return
	}

	// Remove sensitive information like passwords before sending
	for i := range users {
		users[i].Password = ""
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}