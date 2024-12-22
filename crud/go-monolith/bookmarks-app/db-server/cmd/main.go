package main

import (
    "database/sql"
    "log"
    "net/http"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    db, err := sql.Open("sqlite3", "./bookmarks.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Initialize database schema
    createSchema(db)

    // Set up HTTP server
    http.HandleFunc("/api/bookmarks", handleBookmarks(db))
    log.Println("DB Server is running on :8081")
    if err := http.ListenAndServe(":8081", nil); err != nil {
        log.Fatal(err)
    }
}

func createSchema(db *sql.DB) {
    // Create tables for bookmarks, categories, and tags
    createBookmarksTable := `CREATE TABLE IF NOT EXISTS bookmarks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        link TEXT,
        description TEXT,
        category_id INTEGER,
        tag_id INTEGER
    );`
    createCategoriesTable := `CREATE TABLE IF NOT EXISTS categories (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT
    );`
    createTagsTable := `CREATE TABLE IF NOT EXISTS tags (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT
    );`

    statements := []string{createBookmarksTable, createCategoriesTable, createTagsTable}
    for _, statement := range statements {
        _, err := db.Exec(statement)
        if err != nil {
            log.Fatalf("Error creating table: %v", err)
        }
    }
}

func handleBookmarks(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Handle CRUD operations for bookmarks
        // Implementation goes here
    }
}