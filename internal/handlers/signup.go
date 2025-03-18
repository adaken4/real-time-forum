package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"real-time-forum/internal/db"
	"real-time-forum/internal/models"

	"golang.org/x/crypto/bcrypt"
)

// SignupHandler handles user registration
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Bad Request", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Simple validation
	if user.Password != user.PasswordConfirm {
		http.Error(w, "Passwords do not match", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Server Error", err)
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Insert into database
	query := `INSERT INTO users (nickname, first_name, last_name, age, gender, email, password)
	          VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err = db.DB.Exec(query, user.Nickname, user.FirstName, user.LastName, user.Age, user.Gender, user.Email, string(hashedPassword))
	if err != nil {
		fmt.Println("Server Error", err)
		http.Error(w, "Error inserting user", http.StatusInternalServerError)
		return
	}

	// Success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Signup successful!"})
}
