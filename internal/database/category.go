package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, error := c.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)", id, name, description)

	if error != nil {
		return Category{}, error
	}

	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, error := c.db.Query("SELECT id, name, description FROM categories")

	if error != nil {
		return nil, error
	}

	defer rows.Close()

	var categories []Category
	var category Category

	for rows.Next() {
		error := rows.Scan(&category.ID, &category.Name, &category.Description)

		if error != nil {
			return nil, error
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (c *Category) FindByCourseID(courseID string) (Category, error) {
	var category Category

	error := c.db.QueryRow(
		"SELECT c.id, c.name, c.description FROM categories c INNER JOIN courses co ON c.id = co.category_id WHERE co.id = $1",
		courseID).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
	)

	if error != nil {
		return Category{}, error
	}

	return category, nil
}
