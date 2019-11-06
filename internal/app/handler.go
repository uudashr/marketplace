package app

import (
	"encoding/json"
	"errors"
	"time"

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
	b, err := json.Marshal(e.Body)
	if err != nil {
		panic(err)
	}

	h.logger.Log("event", e.Name, "body", string(b), "occuredTime", e.OccuredTime.Format(time.RFC3339))
}
