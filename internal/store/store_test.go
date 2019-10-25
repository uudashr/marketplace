package store_test

import (
	"testing"

	"github.com/shopspring/decimal"

	"github.com/uudashr/marketplace/internal/product"
	"github.com/uudashr/marketplace/internal/store"
)

func TestStore(t *testing.T) {
	cases := map[string]struct {
		id        string
		name      string
		expectErr bool
	}{
		"Default": {
			id:   "an-id",
			name: "SuperMart",
		},

		"Empty ID": {
			id:        "",
			name:      "SuperMart",
			expectErr: true,
		},

		"Empty Name": {
			id:        "an-id",
			name:      "",
			expectErr: true,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			s, err := store.New(c.id, c.name)
			if c.expectErr {
				if err == nil {
					t.Fatal("Expect err")
				}
				return
			}

			if err != nil {
				t.Fatal("err:", err)
			}

			if got, want := s.ID(), c.id; got != want {
				t.Errorf("ID got: %q, want: %q", got, want)
			}

			if got, want := s.Name(), c.name; got != want {
				t.Errorf("Name got: %q, want: %q", got, want)
			}
		})
	}
}

func TestOfferProduct(t *testing.T) {
	cat, err := product.NewCategory(product.NextCategoryID(), "Food")
	if err != nil {
		panic(err)
	}

	cases := map[string]struct {
		id          string
		category    *product.Category
		name        string
		price       decimal.Decimal
		quantity    int
		description string
		expectErr   bool
	}{
		"Default": {
			id:          product.NextID(),
			category:    cat,
			name:        "Mineral Water",
			quantity:    100,
			price:       decimal.NewFromFloat(2500.1),
			description: "Some value",
		},
		"Zero quantity": {
			id:          product.NextID(),
			category:    cat,
			name:        "Mineral Water",
			price:       decimal.NewFromFloat(2500.1),
			description: "Some value",
			quantity:    0,
			expectErr:   true,
		},
		"Negative quantity": {
			id:          product.NextID(),
			category:    cat,
			name:        "Mineral Water",
			price:       decimal.NewFromFloat(2500.1),
			description: "Some value",
			quantity:    -1,
			expectErr:   true,
		},
		"Nil category": {
			id:          product.NextID(),
			category:    nil,
			name:        "Mineral Water",
			quantity:    100,
			price:       decimal.NewFromFloat(2500.1),
			description: "Some value",
			expectErr:   true,
		},
	}

	s, err := store.New(store.NextID(), "My Mart")
	if err != nil {
		t.Fatal("err:", err)
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			p, err := s.OfferProduct(c.id, c.category, c.name, c.price, c.description, c.quantity)
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

			if got, want := p.StoreID(), s.ID(); got != want {
				t.Errorf("StoreID got: %q, want: %q", got, want)
			}

			if got, want := p.CategoryID(), c.category.ID(); got != want {
				t.Errorf("CategoryID got: %q, want: %q", got, want)
			}

			if got, want := p.Name(), c.name; got != want {
				t.Errorf("Name got: %q, want: %q", got, want)
			}

			if got, want := p.Price(), c.price; got != want {
				t.Errorf("Price got: %q, want: %q", got, want)
			}

			if got, want := p.Description(), c.description; got != want {
				t.Errorf("Description got: %q, want: %q", got, want)
			}

			if got, want := p.Quantity(), c.quantity; got != want {
				t.Errorf("Quantity got: %d, want: %d", got, want)
			}
		})
	}
}
