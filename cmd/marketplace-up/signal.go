package main

import (
	"errors"
	"os"
	"os/signal"
	"syscall"
)

func catchSignal(cancel chan struct{}) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	select {
	case <-c:
		return nil
	case <-cancel:
		return errors.New("canceled")
	}
}
