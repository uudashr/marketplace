package app

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/go-kit/kit/log"
	"github.com/uudashr/marketplace/internal/eventd"
)

// LogEventHandler logs all the events.
type LogEventHandler struct {
	logger log.Logger
}

// NewLogEventHandler constructs new LogEventHandler.
func NewLogEventHandler(logger log.Logger) (*LogEventHandler, error) {
	if logger == nil {
		return nil, errors.New("nil logger")
	}

	return &LogEventHandler{
		logger: logger,
	}, nil
}

// HandleEvent implements the eventd.EventHandler interface.
func (h *LogEventHandler) HandleEvent(e eventd.Event) {
	name := reflect.TypeOf(e).Name()
	b, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	h.logger.Log("event", name, "body", string(b))
}
