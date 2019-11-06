package eventd

import (
	"context"
	"reflect"
	"time"
)

// Publisher publishes event.
type Publisher interface {
	Publish(eventName string, eventBody interface{})
}

// Event represent the events.
type Event struct {
	Name        string
	Body        interface{}
	OccuredTime time.Time
}

// EventHandler handles event.
type EventHandler interface {
	HandleEvent(Event)
}

// EventHandlerFunc is the EventHandler function adapter.
type EventHandlerFunc func(Event)

// HandleEvent invokes h(e).
func (h EventHandlerFunc) HandleEvent(e Event) {
	h(e)
}

// Bus represent event bus.
type Bus struct {
	handlers []EventHandler
}

// Publish publishes the event.
func (b *Bus) Publish(eventName string, eventBody interface{}) {
	e := Event{
		Name:        eventName,
		Body:        eventBody,
		OccuredTime: time.Now(),
	}
	for _, h := range b.handlers {
		h.HandleEvent(e)
	}
}

// Subscribe subscribes h to handle events.
func (b *Bus) Subscribe(h EventHandler) {
	for _, v := range b.handlers {
		if v == h {
			// h already subscribed
			return
		}
	}

	b.handlers = append(b.handlers, h)
}

// SubscribeFunc subscribes func as event handler.
func (b *Bus) SubscribeFunc(f func(e Event)) {
	b.Subscribe(EventHandlerFunc(f))
}

type ctxKey int

const (
	pubCtxKey ctxKey = iota + 1
)

// ContextWithPublisher wrap context with publisher.
func ContextWithPublisher(ctx context.Context, p Publisher) context.Context {
	return context.WithValue(ctx, pubCtxKey, p)
}

// PublisherFromContext get publisher from context if exist.
func PublisherFromContext(ctx context.Context) Publisher {
	v := ctx.Value(pubCtxKey)
	p, ok := v.(Publisher)
	if !ok {
		return nil
	}

	return p
}

// Publish publishs event using publiher on the ctx.
func Publish(ctx context.Context, eventName string, eventBody interface{}) {
	pub := PublisherFromContext(ctx)
	if pub == nil {
		return
	}

	pub.Publish(eventName, eventBody)
}

// PublishNamed publishes event using publisher on the ctx.
// Name of the event will be derinfed from the type name of eventBody.
func PublishNamed(ctx context.Context, eventBody interface{}) {
	eventName := reflect.TypeOf(eventBody).Name()
	Publish(ctx, eventName, eventBody)
}
