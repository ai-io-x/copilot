package database

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

const dbFile = "websites.db"

// InitializeDatabase creates the SQLite database and the necessary tables.
func InitializeDatabase() *sql.DB {
    db, err := sql.Open("sqlite3", dbFile)
    if err != nil {
        log.Fatalf("failed to open database: %v", err)
    }

    createTables(db)

    return db
}

// createTables creates the necessary tables in the database.
func createTables(db *sql.DB) {
    createWebsitesTable := `CREATE TABLE IF NOT EXISTS websites (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        link TEXT NOT NULL,
        description TEXT,
        category_id INTEGER,
        FOREIGN KEY (category_id) REFERENCES categories(id)
    );`

    createCategoriesTable := `CREATE TABLE IF NOT EXISTS categories (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL
    );`

    createTagsTable := `CREATE TABLE IF NOT EXISTS tags (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL
    );`

    statements := []string{createWebsitesTable, createCategoriesTable, createTagsTable}

    for _, statement := range statements {
        _, err := db.Exec(statement)
        if err != nil {
            log.Fatalf("failed to create table: %v", err)
        }
    }
}

// CloseDatabase closes the database connection.
func CloseDatabase(db *sql.DB) {
    if err := db.Close(); err != nil {
        log.Fatalf("failed to close database: %v", err)
    }
}