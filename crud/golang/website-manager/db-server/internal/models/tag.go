type Tag struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func (t *Tag) Create(db *sql.DB) error {
    query := "INSERT INTO tags (name) VALUES (?)"
    _, err := db.Exec(query, t.Name)
    return err
}

func (t *Tag) Update(db *sql.DB) error {
    query := "UPDATE tags SET name = ? WHERE id = ?"
    _, err := db.Exec(query, t.Name, t.ID)
    return err
}

func (t *Tag) Delete(db *sql.DB) error {
    query := "DELETE FROM tags WHERE id = ?"
    _, err := db.Exec(query, t.ID)
    return err
}

func GetTags(db *sql.DB) ([]Tag, error) {
    query := "SELECT id, name FROM tags"
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tags []Tag
    for rows.Next() {
        var tag Tag
        if err := rows.Scan(&tag.ID, &tag.Name); err != nil {
            return nil, err
        }
        tags = append(tags, tag)
    }
    return tags, nil
}