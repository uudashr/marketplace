package mysql

import (
	"database/sql"
	"errors"

	"github.com/uudashr/marketplace/internal/store"
)

// StoreRepository is repository for Store.
type StoreRepository struct {
	db *sql.DB
}

// NewStoreRepository constructs new store repository.
func NewStoreRepository(db *sql.DB) (*StoreRepository, error) {
	return &StoreRepository{
		db: db,
	}, nil
}

// Store stores/puts store to the repository.
func (r *StoreRepository) Store(str *store.Store) error {
	res, err := r.db.Exec("INSERT INTO stores (id, name) VALUES (?, ?)", str.ID(), str.Name())
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

// StoreByID retrieve store by ID.
func (r *StoreRepository) StoreByID(id string) (*store.Store, error) {
	var (
		name string
	)
	err := r.db.QueryRow("SELECT name FROM stores WHERE id = ?", id).Scan(&name)
	if err != nil {
		return nil, err
	}

	return store.New(id, name)
}

// Stores retrieves stores.
func (r *StoreRepository) Stores() ([]*store.Store, error) {
	rows, err := r.db.Query("SELECT id, name FROM stores")
	if err != nil {
		return nil, err
	}

	var out []*store.Store
	for rows.Next() {
		var (
			id   string
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}

		str, err := store.New(id, name)
		if err != nil {
			return nil, err
		}

		out = append(out, str)
	}

	return out, nil
}
