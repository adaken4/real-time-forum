package main

import (
	"log"
	"net/http"
	"real-time-forum/internal/auth"
	"real-time-forum/internal/db"
	"real-time-forum/internal/handlers"
)

func main() {
	// Initialize the database
	db.Init()
	defer db.CloseDB()

	mux := http.NewServeMux()

	// Serve static files
	fs := http.FileServer(http.Dir("frontend/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.Handle("/ws", auth.SessionMiddleware(http.HandlerFunc(handlers.WebSocketHandler)))
	mux.HandleFunc("/api/signup", handlers.SignupHandler)
	mux.HandleFunc("/api/signin", handlers.SigninHandler)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/index.html")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("Server listening on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
