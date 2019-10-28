package fixture

import (
	"math/rand"

	"github.com/icrowley/fake"
	"github.com/shopspring/decimal"
	"github.com/uudashr/marketplace/internal/product"
	"github.com/uudashr/marketplace/internal/store"
)

// ProductOfStore fixture of products of a store.
func ProductOfStore(str *store.Store) *product.Product {
	return ProductOfStoreWithOptions(str, ProductOptions{})
}

// Product fixture.
func Product() *product.Product {
	str := Store()
	return ProductOfStore(str)
}

// ProductOfStoreWithOptions fixture of products of a store with specific options.
func ProductOfStoreWithOptions(str *store.Store, opts ProductOptions) *product.Product {
	prd, err := product.New(
		product.NextID(),
		str.ID(),
		product.NextCategoryID(),
		opts.nameOption(),
		opts.priceOption(),
		opts.descriptionOption(),
		opts.quantityOption())
	if err != nil {
		panic(err)
	}

	return prd
}

// ProductsOfStore fixture.
func ProductsOfStore(str *store.Store, n int) []*product.Product {
	out := make([]*product.Product, n)
	for i := 0; i < n; i++ {
		out[i] = ProductOfStore(str)
	}
	return out
}

// Products fixture.
func Products(n int) []*product.Product {
	out := make([]*product.Product, n)
	for i := 0; i < n; i++ {
		out[i] = Product()
	}
	return out
}

// ProductOptions is product options.
type ProductOptions struct {
	Name string
}

func (opts ProductOptions) nameOption() string {
	if opts.Name == "" {
		return fake.ProductName()
	}

	return opts.Name
}

func (opts ProductOptions) priceOption() decimal.Decimal {
	return decimal.NewFromFloat(float64(2500 + rand.Intn(297_501)))
}

func (opts ProductOptions) descriptionOption() string {
	return fake.Paragraph()
}

func (opts ProductOptions) quantityOption() int {
	return 100 + rand.Intn(201)
}
