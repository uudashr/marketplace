package fixture

import (
	"github.com/icrowley/fake"
	"github.com/uudashr/marketplace/internal/store"
)

// Store fixture.
func Store() *store.Store {
	str, err := store.New(store.NextID(), fake.Company())
	if err != nil {
		panic(err)
	}
	return str
}

// Stores fixture.
func Stores(n int) []*store.Store {
	out := make([]*store.Store, n)
	for i := 0; i < n; i++ {
		out[i] = Store()
	}
	return out
}
