package mysql

import (
	"database/sql"
	"errors"

	"github.com/uudashr/marketplace/internal/category"
)

// CategoryRepository is repository for Category.
type CategoryRepository struct {
	db *sql.DB
}

// NewCategoryRepository constructs new category repository.
func NewCategoryRepository(db *sql.DB) (*CategoryRepository, error) {
	return &CategoryRepository{
		db: db,
	}, nil
}

// Store the category to the repository.
func (r *CategoryRepository) Store(cat *category.Category) error {
	res, err := r.db.Exec("INSERT INTO categories (id, name) VALUES (?, ?)", cat.ID(), cat.Name())
	// TODO: how to handle unique name -> Error 1062: Duplicate entry
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

// CategoryByID on the repository.
func (r *CategoryRepository) CategoryByID(id string) (*category.Category, error) {
	var (
		name string
	)
	err := r.db.QueryRow("SELECT name FROM categories WHERE id = ?", id).Scan(&name)
	if err != nil {
		return nil, err
	}

	return category.New(id, name)
}

// Categories on the repository.
func (r *CategoryRepository) Categories() ([]*category.Category, error) {
	rows, err := r.db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}

	var out []*category.Category
	for rows.Next() {
		var (
			id   string
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}

		cat, err := category.New(id, name)
		if err != nil {
			return nil, err
		}

		out = append(out, cat)
	}

	return out, nil
}
