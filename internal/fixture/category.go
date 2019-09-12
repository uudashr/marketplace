package fixture

import (
	"github.com/icrowley/fake"
	"github.com/uudashr/marketplace/internal/product"
)

// Category fixture.
func Category() *product.Category {
	cat, err := product.NewCategory(product.NextCategoryID(), fake.Industry())
	if err != nil {
		panic(err)
	}
	return cat
}

// Categories fixture.
func Categories(n int) []*product.Category {
	out := make([]*product.Category, n)
	for i := 0; i < n; i++ {
		out[i] = Category()
	}
	return out
}
