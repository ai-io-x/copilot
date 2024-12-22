package model

type Bookmark struct {
    ID          int64    `json:"id" db:"id"`
    Name        string   `json:"name" db:"name"`
    URL         string   `json:"url" db:"url"`
    Description string   `json:"description" db:"description"`
    Categories  []string `json:"categories"`
    Tags        []string `json:"tags"`
}