package repository

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

type SQLiteRepository struct {
    db *sql.DB
}

func NewSQLiteRepository(dataSourceName string) (*SQLiteRepository, error) {
    db, err := sql.Open("sqlite3", dataSourceName)
    if err != nil {
        return nil, err
    }

    return &SQLiteRepository{db: db}, nil
}

func (r *SQLiteRepository) CreateBookmark(name string, link string) error {
    _, err := r.db.Exec("INSERT INTO bookmarks (name, link) VALUES (?, ?)", name, link)
    return err
}

func (r *SQLiteRepository) GetBookmarks() ([]Bookmark, error) {
    rows, err := r.db.Query("SELECT id, name, link FROM bookmarks")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var bookmarks []Bookmark
    for rows.Next() {
        var b Bookmark
        if err := rows.Scan(&b.ID, &b.Name, &b.Link); err != nil {
            return nil, err
        }
        bookmarks = append(bookmarks, b)
    }
    return bookmarks, nil
}

func (r *SQLiteRepository) UpdateBookmark(id int, name string, link string) error {
    _, err := r.db.Exec("UPDATE bookmarks SET name = ?, link = ? WHERE id = ?", name, link, id)
    return err
}

func (r *SQLiteRepository) DeleteBookmark(id int) error {
    _, err := r.db.Exec("DELETE FROM bookmarks WHERE id = ?", id)
    return err
}

func (r *SQLiteRepository) Close() error {
    return r.db.Close()
}

type Bookmark struct {
    ID   int
    Name string
    Link string
}