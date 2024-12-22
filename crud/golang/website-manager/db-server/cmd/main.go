package main

import (
    "database/sql"
    "log"
    "net/http"

    _ "github.com/mattn/go-sqlite3"
)

func main() {
    db, err := sql.Open("sqlite3", "./website_manager.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Initialize the database and create tables if they don't exist
    createTables(db)

    // Start the HTTP server
    http.HandleFunc("/api/websites", handleWebsites(db))
    log.Println("Database server is running on port 8081...")
    if err := http.ListenAndServe(":8081", nil); err != nil {
        log.Fatal(err)
    }
}

func createTables(db *sql.DB) {
    // Create tables for websites, categories, and tags
    createWebsitesTable := `CREATE TABLE IF NOT EXISTS websites (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        link TEXT NOT NULL,
        description TEXT
    );`
    createCategoriesTable := `CREATE TABLE IF NOT EXISTS categories (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL
    );`
    createTagsTable := `CREATE TABLE IF NOT EXISTS tags (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL
    );`

    if _, err := db.Exec(createWebsitesTable); err != nil {
        log.Fatal(err)
    }
    if _, err := db.Exec(createCategoriesTable); err != nil {
        log.Fatal(err)
    }
    if _, err := db.Exec(createTagsTable); err != nil {
        log.Fatal(err)
    }

    // Insert example data
    insertExampleData(db)
}

func insertExampleData(db *sql.DB) {
    // Insert example data into the websites table
    exampleWebsites := []struct {
        name        string
        link        string
        description string
    }{
        {"Example Site 1", "https://example1.com", "This is an example website."},
        {"Example Site 2", "https://example2.com", "This is another example website."},
        {"Example Site 3", "https://example3.com", "Yet another example website."},
    }

    for _, website := range exampleWebsites {
        _, err := db.Exec("INSERT INTO websites (name, link, description) VALUES (?, ?, ?)", website.name, website.link, website.description)
        if err != nil {
            log.Fatal(err)
        }
    }
}

// handleWebsites handles HTTP requests for website data
func handleWebsites(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Handle CRUD operations here
        // This is a placeholder for the actual implementation
        w.Write([]byte("Handle CRUD operations for websites"))
    }
}