package fixture

import (
	"math/rand"

	"github.com/icrowley/fake"
	"github.com/shopspring/decimal"
	"github.com/uudashr/marketplace/internal/product"
	"github.com/uudashr/marketplace/internal/store"
)

func ProductOfStore(str *store.Store) *product.Product {
	return ProductOfStoreWithOptions(str, ProductOptions{})
}

func ProductOfStoreWithOptions(str *store.Store, opts ProductOptions) *product.Product {
	prod, err := product.New(
		product.NextID(),
		str.ID(),
		product.NextCategoryID(),
		opts.nameOption(),
		opts.priceOption(),
		opts.quantityOption(),
		opts.descriptionOption())
	if err != nil {
		panic(err)
	}

	return prod
}

type ProductOptions struct {
	Name string
}

func (opts ProductOptions) nameOption() string {
	if opts.Name == "" {
		return fake.ProductName()
	}

	return opts.Name
}

func (opts ProductOptions) descriptionOption() string {
	return fake.Paragraph()
}

func (opts ProductOptions) quantityOption() int {
	return 100 + rand.Intn(201)
}

func (opts ProductOptions) priceOption() decimal.Decimal {
	return decimal.NewFromFloat(float64(2500 + rand.Intn(297_501)))
}
