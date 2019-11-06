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

	eventName, eventBody := "Hi", "Hello, World!"
	b.Publish(eventName, eventBody)

	if got, want := captured.Name, eventName; got != want {
		t.Errorf("Captured event name %q, want: %q", got, want)
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
	eventd.PublishNamed(ctx, msg)

	if got, want := captured.Name, "Ping"; got != want {
		t.Errorf("Captured event name %q, want: %q", got, want)
	}
}

type Ping struct {
	Message string
}
