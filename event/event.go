package event

import (
	"github.com/shopspring/decimal"
)

// NewProductCreated is events raised when new product created.
type NewProductCreated struct {
	ID          string
	StoreID     string
	CategoryID  string
	Name        string
	Price       decimal.Decimal
	Description string
	Quantity    int
}
