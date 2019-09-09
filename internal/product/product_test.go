package product_test

import (
	"testing"

	"github.com/uudashr/marketplace/internal/category"

	"github.com/uudashr/marketplace/internal/product"
	"github.com/uudashr/marketplace/internal/store"

	"github.com/shopspring/decimal"
)

func TestProduct(t *testing.T) {
	cases := map[string]struct {
		id          string
		storeID     string
		categoryID  string
		name        string
		price       decimal.Decimal
		quantity    int
		description string
		expectErr   bool
	}{
		"Default": {
			id:          product.NextID(),
			storeID:     store.NextID(),
			categoryID:  category.NextID(),
			name:        "Mineral Water",
			price:       decimal.NewFromFloat(2500.1),
			quantity:    100,
			description: "Some value",
		},
		"Zero quantity": {
			id:          product.NextID(),
			storeID:     store.NextID(),
			categoryID:  category.NextID(),
			name:        "Mineral Water",
			price:       decimal.NewFromFloat(2500.1),
			quantity:    0,
			description: "Some value",
		},
		"Empty description": {
			id:          product.NextID(),
			storeID:     store.NextID(),
			categoryID:  category.NextID(),
			name:        "Mineral Water",
			price:       decimal.NewFromFloat(2500.1),
			quantity:    100,
			description: "",
		},
		"Empty name": {
			id:          product.NextID(),
			storeID:     store.NextID(),
			categoryID:  category.NextID(),
			name:        "",
			price:       decimal.NewFromFloat(2500.1),
			quantity:    100,
			description: "Some value",
			expectErr:   true,
		},
		"Negative price": {
			id:          product.NextID(),
			storeID:     store.NextID(),
			categoryID:  category.NextID(),
			name:        "Mineral Water",
			price:       decimal.NewFromFloat(-2500.1),
			quantity:    100,
			description: "Some value",
			expectErr:   true,
		},
		"Negative quantity": {
			id:          product.NextID(),
			storeID:     store.NextID(),
			categoryID:  category.NextID(),
			name:        "Mineral Water",
			price:       decimal.NewFromFloat(2500.1),
			quantity:    -10,
			description: "Some value",
			expectErr:   true,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			p, err := product.New(c.id, c.storeID, c.categoryID, c.name, c.price, c.quantity, c.description)
			if c.expectErr {
				if err == nil {
					t.Fatal("Expect err")
				}
				return
			}

			if err != nil {
				t.Fatal("err:", err)
			}

			if got, want := p.ID(), c.id; got != want {
				t.Errorf("ID got: %q, want: %q", got, want)
			}

			if got, want := p.StoreID(), c.storeID; got != want {
				t.Errorf("StoreID got: %q, want: %q", got, want)
			}

			if got, want := p.CategoryID(), c.categoryID; got != want {
				t.Errorf("CategoryID got: %q, want: %q", got, want)
			}

			if got, want := p.Name(), c.name; got != want {
				t.Errorf("Name got: %q, want: %q", got, want)
			}

			if got, want := p.Price(), c.price; got != want {
				t.Errorf("Price got: %q, want: %q", got, want)
			}

			if got, want := p.Quantity(), c.quantity; got != want {
				t.Errorf("Quantity got: %d, want: %d", got, want)
			}

			if got, want := p.Description(), c.description; got != want {
				t.Errorf("Description got: %q, want: %q", got, want)
			}
		})
	}
}
