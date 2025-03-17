package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite3", "./rt-forum.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	createTables()
	createCategories()
}

func createTables() {
	sqlBytes, err := os.ReadFile("internal/db/schema.sql")
	if err != nil {
		log.Fatalf("Failed to read schema file: %v\n", err)
	}

	sqlStatements := string(sqlBytes)

	if _, err := DB.Exec(sqlStatements); err != nil {
		log.Fatalf("Failed to execute statements: %v\nQuery %s\n", err, sqlStatements)
	}

	log.Println("All tables created successfully.")
}

func createCategories() {
	stmt, err := DB.Prepare("INSERT OR IGNORE INTO categories (name, description) VALUES (?, ?)")
	if err != nil {
		log.Fatalf("Failed to prepare category insert: %v", err)
	}
	defer stmt.Close()
	
	predefinedCategories := []struct {
		Name        string
		Description string
	}{
		{"Technology", "Posts related to the latest technology and trends"},
		{"Health", "Discussions about health, fitness, and well-being"},
		{"Education", "Topics about learning and education"},
		{"Entertainment", "Movies, music, games, and all things fun"},
		{"Lifestyle", "Fashion, home decor, and daily living tips"},
		{"Travel", "Exploring the world, sharing travel experiences"},
	}

	for _, category := range predefinedCategories {
		_, err := stmt.Exec(category.Name, category.Description)
		if err != nil {
			log.Printf("Error inserting category '%s': '%v'", category.Name, err)
		}
	}
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
