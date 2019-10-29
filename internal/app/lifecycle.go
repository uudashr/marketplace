package app

import (
	"context"

	"github.com/uudashr/marketplace/internal/eventd"
)

// Lifecycle represents applicataion lifecycle.
// This will capture the events by providing event publisher via context.Context.
type Lifecycle struct {
	bus      eventd.Bus
	events   []eventd.Event
	handlers []eventd.EventHandler
}

// NewLifecycle construct new lifecycle.
func NewLifecycle(ctx context.Context) (*Lifecycle, context.Context) {
	lc := new(Lifecycle)

	lc.bus.SubscribeFunc(func(e eventd.Event) {
		lc.events = append(lc.events, e)
	})

	pubCtx := eventd.ContextWithPublisher(ctx, &lc.bus)
	return lc, pubCtx
}

// End ends the lifecycle.
// if err nill then it will process the events on the Handler.
func (lc *Lifecycle) End(err error) {
	if err != nil {
		return
	}

	for _, e := range lc.events {
		handleEvent(lc.handlers, e)
	}
}

// SubscribeEvent subscribes for events.
func (lc *Lifecycle) SubscribeEvent(h eventd.EventHandler) {
	for _, v := range lc.handlers {
		if v == h {
			// h already subscribed
			return
		}
	}

	lc.handlers = append(lc.handlers, h)
}

// SubscribeEventFunc subscribes for events using func.
func (lc *Lifecycle) SubscribeEventFunc(f func(e eventd.Event)) {
	lc.SubscribeEvent(eventd.EventHandlerFunc(f))
}

func handleEvent(hs []eventd.EventHandler, e eventd.Event) {
	for _, h := range hs {
		h.HandleEvent(e)
	}
}
