package main

import (
	"log"
	"net/http"
	"real-time-forum/internal/db"
)

func main() {
	// Initialize the database
	db.Init()
	defer db.CloseDB()

	mux := http.NewServeMux()

	// Serve static files
	fs := http.FileServer(http.Dir("frontend/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/index.html")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
