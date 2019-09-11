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
