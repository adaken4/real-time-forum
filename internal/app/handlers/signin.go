package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"real-time-forum/internal/db"
	"real-time-forum/internal/models"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func SigninHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var creds models.User
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Retrieve user from database
	var userID int
	var storedPassword string
	query := `SELECT user_id, password FROM users WHERE email = ? OR nickname = ?`
	err = db.DB.QueryRow(query, creds.Email, creds.Nickname).Scan(&userID, &storedPassword)
	if err != nil {
		log.Printf("Database query error: %v", err)
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(creds.Password))
	if err != nil {
		http.Error(w, "Invalid email/username or password", http.StatusUnauthorized)
		return
	}

	// Delete any existing session for the user (enforcing single-session authentication)
	deleteQuery := `DELETE FROM sessions WHERE user_id = ?`
	_, err = db.DB.Exec(deleteQuery, userID)
	if err != nil {
		http.Error(w, "Failed to clear old sessions", http.StatusInternalServerError)
		log.Printf("Database delete error: %v", err)
		return
	}

	// Create a session
	sessionID := uuid.New().String()
	expiration := time.Now().Add(24 * time.Hour) // 1-day session expiration
	insertQuery := `INSERT INTO sessions (session_id, user_id, expires_at) VALUES (?, ?, ?)`
	_, err = db.DB.Exec(insertQuery, sessionID, userID, expiration)
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		log.Printf("Database insert error: %v", err)
		return
	}

	// Set a session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Expires:  expiration,
		Path:     "/",
		HttpOnly: true, // Prevent JavaScript access
		SameSite: http.SameSiteLaxMode,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Signin successful!"})

}
