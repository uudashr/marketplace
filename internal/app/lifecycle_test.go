package app_test

import (
	"context"
	"errors"
	"testing"

	"github.com/uudashr/marketplace/internal/eventd"

	"github.com/uudashr/marketplace/internal/app"
)

func TestLifecycle_Publish(t *testing.T) {
	cases := map[string]struct {
		event         interface{}
		err           error
		expectHandled bool
	}{
		"No Error": {
			event:         Ping{Message: "Hello"},
			err:           nil,
			expectHandled: true,
		},
		"Has Error": {
			event:         Ping{Message: "World"},
			err:           errors.New("Opps"),
			expectHandled: false,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			lc, ctx := app.NewLifecycle(context.TODO())
			var handled bool
			lc.SubscribeEventFunc(func(e eventd.Event) {
				if e.Body == c.event {
					handled = true
				}
			})

			eventd.PublishNamed(ctx, c.event)

			lc.End(c.err)

			if got, want := handled, c.expectHandled; got != want {
				t.Errorf("Handled got: %t, want: %t", got, want)
			}
		})
	}

}

type Ping struct {
	Message string
}
