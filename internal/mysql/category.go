package mysql

import (
	"database/sql"
	"errors"

	"github.com/uudashr/marketplace/internal/product"
)

// CategoryRepository is repository for product category.
type CategoryRepository struct {
	db *sql.DB
}

// NewCategoryRepository constructs new product category repository.
func NewCategoryRepository(db *sql.DB) (*CategoryRepository, error) {
	return &CategoryRepository{
		db: db,
	}, nil
}

// Store stores the product category to the repository.
func (r *CategoryRepository) Store(cat *product.Category) error {
	res, err := r.db.Exec("INSERT INTO categories (id, name) VALUES (?, ?)", cat.ID(), cat.Name())
	// TODO: how to handle unique name -> Error 1062: Duplicate entry
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

// CategoryByID retrieves product category by id.
func (r *CategoryRepository) CategoryByID(id string) (*product.Category, error) {
	var (
		name string
	)
	err := r.db.QueryRow("SELECT name FROM categories WHERE id = ?", id).Scan(&name)
	if err != nil {
		return nil, err
	}

	return product.NewCategory(id, name)
}

// Categories retrieves product categories.
func (r *CategoryRepository) Categories() ([]*product.Category, error) {
	rows, err := r.db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}

	var out []*product.Category
	for rows.Next() {
		var (
			id   string
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}

		cat, err := product.NewCategory(id, name)
		if err != nil {
			return nil, err
		}

		out = append(out, cat)
	}

	return out, nil
}
