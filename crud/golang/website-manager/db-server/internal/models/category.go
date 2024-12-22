package models

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c *Category) CreateCategory(db *sql.DB) error {
	// Implementation for creating a category in the database
}

func (c *Category) GetCategory(db *sql.DB, id int) (*Category, error) {
	// Implementation for fetching a category by ID from the database
}

func (c *Category) UpdateCategory(db *sql.DB) error {
	// Implementation for updating a category in the database
}

func (c *Category) DeleteCategory(db *sql.DB, id int) error {
	// Implementation for deleting a category from the database
}

func GetAllCategories(db *sql.DB) ([]Category, error) {
	// Implementation for fetching all categories from the database
}