package fixture

import (
	"github.com/icrowley/fake"

	"github.com/uudashr/marketplace/internal/category"
)

// Category fixture.
func Category() *category.Category {
	cat, err := category.New(category.NextID(), fake.Industry())
	if err != nil {
		panic(err)
	}
	return cat
}

// Categories fixture.
func Categories(n int) []*category.Category {
	out := make([]*category.Category, n)
	for i := 0; i < n; i++ {
		out[i] = Category()
	}
	return out
}
