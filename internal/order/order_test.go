package order_test

import (
	"testing"
	"time"

	"github.com/uudashr/marketplace/internal/product"

	"github.com/shopspring/decimal"

	"github.com/uudashr/marketplace/internal/order"
)

func TestItem(t *testing.T) {
	productID := product.NextID()
	price := decimal.NewFromFloat(2500.0)
	quantity := 100
	it, err := order.NewItem(productID, price, quantity)
	if err != nil {
		t.Fatal("err:", err)
	}

	if got, want := it.ProductID(), productID; got != want {
		t.Errorf("ProductID got: %q, want: %q", got, want)
	}

	if got, want := it.Price(), price; !got.Equal(want) {
		t.Errorf("Price got: %q, want: %q", got, want)
	}

	if got, want := it.Quantity(), quantity; got != want {
		t.Errorf("ProductID got: %d, want: %d", got, want)
	}
}

func TestOrder(t *testing.T) {
	it1, err := order.NewItem(product.NextID(), decimal.NewFromFloat(2500.0), 10)
	if err != nil {
		t.Fatal("err:", err)
	}

	it2, err := order.NewItem(product.NextID(), decimal.NewFromFloat(2500.0), 10)
	if err != nil {
		t.Fatal("err:", err)
	}

	items := []*order.Item{it1, it2}
	ord, err := order.New(order.NextID(), items, order.Initiated, time.Now())
	if err != nil {
		t.Fatal("err:", err)
	}

	_ = ord
}
