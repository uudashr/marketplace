package main

import (
	"errors"
	"io"

	"github.com/go-kit/kit/log"
)

func logWithFormat(writer io.Writer, format string) (log.Logger, error) {
	switch format {
	case "logfmt":
		return log.NewLogfmtLogger(writer), nil
	case "json":
		return log.NewJSONLogger(writer), nil
	default:
		return nil, errors.New("unrecognized format")
	}
}
