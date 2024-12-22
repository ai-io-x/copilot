type Website struct {
    ID          int      `json:"id"`
    Name        string   `json:"name"`
    Link        string   `json:"link"`
    Description string   `json:"description"`
    Tags        []string `json:"tags"`
    Category    string   `json:"category"`
}

func (w *Website) Validate() error {
    if w.Name == "" {
        return errors.New("website name is required")
    }
    if w.Link == "" {
        return errors.New("website link is required")
    }
    return nil
}