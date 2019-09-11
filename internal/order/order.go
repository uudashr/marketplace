package order

import (
	"errors"
	"time"

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

// Fulfillment represents the fulfillment.
type Fulfillment struct {
	id          string
	items       []*Item
	status      Status
	createdTime time.Time
}

// New constructs new Fulfillment instance.
func New(id string, items []*Item, status Status, createdTime time.Time) (*Fulfillment, error) {
	if id == "" {
		return nil, errors.New("empty id")
	}

	if len(items) == 0 {
		return nil, errors.New("empty items")
	}

	if createdTime.IsZero() {
		return nil, errors.New("zero createdTime")
	}

	return &Fulfillment{
		id:          id,
		items:       items,
		status:      status,
		createdTime: createdTime,
	}, nil
}

// ID of the fulfillment.
func (o Fulfillment) ID() string {
	return o.id
}

// Items of the fulfillment.
func (o Fulfillment) Items() []*Item {
	return o.items
}

// Status of the fulfillment.
func (o Fulfillment) Status() Status {
	return o.status
}

// CreatedTime of the fulfillment.
func (o Fulfillment) CreatedTime() time.Time {
	return o.createdTime
}

// NextID returns unique id for Order.
func NextID() string {
	return xid.New().String()
}

// Item represents the fulfillment item (or line item).
type Item struct {
	storeProductID string
	price          decimal.Decimal
	quantity       int
}

// NewItem constructs new item.
func NewItem(storeProductID string, price decimal.Decimal, quantity int) (*Item, error) {
	if storeProductID == "" {
		return nil, errors.New("empty storeProductID")
	}

	if !price.GreaterThan(decimal.Zero) {
		return nil, errors.New("zero or negative price")
	}

	if quantity == 0 {
		return nil, errors.New("zero quantity")
	}

	return &Item{
		storeProductID: storeProductID,
		price:          price,
		quantity:       quantity,
	}, nil
}

// StoreProductID of the item.
func (i Item) StoreProductID() string {
	return i.storeProductID
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
