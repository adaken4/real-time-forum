package auth

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"real-time-forum/internal/db"
)

type contextKey string

const userIDKey contextKey = "userID"

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil || cookie.Value == "" {
			// No session cookie, proceed without user context
			next.ServeHTTP(w, r)
			return
		}

		// Validate the session
		var userID string
		var expiresAt time.Time
		query := `SELECT user_id, expires_at FROM sessions WHERE session_id = ?`
		err = db.DB.QueryRow(query, cookie.Value).Scan(&userID, &expiresAt)
		if err == sql.ErrNoRows || time.Now().After(expiresAt) {
			// Invalid or expired session, clear the cookie
			http.SetCookie(w, &http.Cookie{
				Name:     "session_id",
				Value:    "",
				Expires:  time.Now().Add(-time.Hour),
				Path:     "/",
				HttpOnly: true,
			})
			next.ServeHTTP(w, r)
			return
		}

		// Add userID to the request context
		ctx := context.WithValue(r.Context(), userIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RedirectIfAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(userIDKey)
		if userID != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetUserID(r *http.Request) (string, bool) {
	userID, ok := r.Context().Value(userIDKey).(string)
	return userID, ok
}

func IsAuthenticated(r *http.Request) bool {
	_, ok := GetUserID(r)
	return ok
}

func UserIDKey() contextKey {
	return userIDKey
}

type AuthStatus struct {
	Authenticated bool   `json:"authenticated"`
	UserID        string `json:"user_id,omitempty"`
}

func AuthStatusHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := GetUserID(r)
	response := AuthStatus{Authenticated: ok}

	if ok {
		response.UserID = userID
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
