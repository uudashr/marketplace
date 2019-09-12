package order

import (
	"errors"
	"time"

	"github.com/uudashr/marketplace/internal/product"

	"github.com/rs/xid"

	"github.com/shopspring/decimal"
)

// Status represent status of the order.
type Status int

const (
	// Initiated means order ready to pick by store.
	Initiated Status = iota

	// Confirmed means order are confirmed up by the store and will be sent soon.
	Confirmed

	// Sent means the items are sent by store (in delivery).
	Sent

	// Delivered means the items has been delivered to destination.
	Delivered

	// Canceled means the order canceled by user or system.
	Canceled

	// Rejected means the order are canceled by by the store.
	Rejected
)

// Order represents the order.
type Order struct {
	id          string
	items       []*Item
	status      Status
	createdTime time.Time
}

// New constructs new order instance.
func New(id string, items []*Item, status Status, createdTime time.Time) (*Order, error) {
	if id == "" {
		return nil, errors.New("empty id")
	}

	if len(items) == 0 {
		return nil, errors.New("empty items")
	}

	if createdTime.IsZero() {
		return nil, errors.New("zero createdTime")
	}

	return &Order{
		id:          id,
		items:       items,
		status:      status,
		createdTime: createdTime,
	}, nil
}

// ID of the order.
func (o Order) ID() string {
	return o.id
}

// Items of the order.
func (o Order) Items() []*Item {
	return o.items
}

// Status of the order.
func (o Order) Status() Status {
	return o.status
}

// CreatedTime of the order.
func (o Order) CreatedTime() time.Time {
	return o.createdTime
}

// NextID returns unique id for order.
func NextID() string {
	return xid.New().String()
}

// Item represents the order item (or line item).
type Item struct {
	productID string
	price     decimal.Decimal
	quantity  int
}

// NewItem constructs new item.
func NewItem(productID string, price decimal.Decimal, quantity int) (*Item, error) {
	if productID == "" {
		return nil, errors.New("empty productID")
	}

	if !price.GreaterThan(decimal.Zero) {
		return nil, errors.New("zero or negative price")
	}

	if quantity == 0 {
		return nil, errors.New("zero quantity")
	}

	return &Item{
		productID: productID,
		price:     price,
		quantity:  quantity,
	}, nil
}

// ProductID of the item.
func (i Item) ProductID() string {
	return i.productID
}

// Price of the item.
func (i Item) Price() decimal.Decimal {
	return i.price
}

// Quantity of the item.
func (i Item) Quantity() int {
	return i.quantity
}

// SubTotal price of the item.
func (i Item) SubTotal() decimal.Decimal {
	return i.price.Mul(decimal.NewFromFloat(float64(i.quantity)))
}
