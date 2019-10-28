package eventd_test

import (
	"context"
	"testing"

	"github.com/uudashr/marketplace/internal/eventd"
)

func TestBus(t *testing.T) {
	var b eventd.Bus

	var captured eventd.Event
	b.SubscribeFunc(func(e eventd.Event) {
		captured = e
	})

	msg := "Hello"
	b.Publish(msg)

	if got, want := captured, msg; got != want {
		t.Errorf("Captured %q, want: %q", got, want)
	}
}

func TestPublisherFromContext(t *testing.T) {
	b := new(eventd.Bus)

	ctx := eventd.ContextWithPublisher(context.TODO(), b)
	pub := eventd.PublisherFromContext(ctx)
	if got, want := pub, b; got != want {
		t.Fatalf("Publisher got: %p, want: %p", got, want)
	}
}

func TestPublish(t *testing.T) {
	var b eventd.Bus
	ctx := eventd.ContextWithPublisher(context.TODO(), &b)

	var captured eventd.Event
	b.SubscribeFunc(func(e eventd.Event) {
		captured = e
	})

	msg := Ping{
		Message: "Hello",
	}
	eventd.Publish(ctx, msg)

	if got, want := captured, msg; got != want {
		t.Errorf("Captured %q, want: %q", got, want)
	}
}

type Ping struct {
	Message string
}
