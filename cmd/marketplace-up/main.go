package main

import (
	"context"
	nethttp "net/http"
	"os"
	"time"

	"github.com/uudashr/marketplace/internal/app"
	"github.com/uudashr/marketplace/internal/inmem"

	"github.com/uudashr/marketplace/internal/http"

	"github.com/go-kit/kit/log"

	"github.com/oklog/run"
)

func main() {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	mainLogger := log.With(logger, "component", "Main")

	var g run.Group

	mainLogger.Log("msg", "Starting marketplace")

	// Signal catcher
	{
		cancel := make(chan struct{})
		catchLogger := log.With(logger, "component", "SignalCatcher")
		g.Add(func() error {
			err := catchSignal(cancel)
			if err != nil {
				return err
			}

			catchLogger.Log("msg", "Got shutdown signal")
			return nil
		}, func(error) {
			close(cancel)
		})
	}

	// HTTP
	{
		httpLogger := log.With(logger, "component", "HTTP")
		categoryRepo := inmem.NewCategoryRepository()
		appService, err := app.NewService(categoryRepo)
		if err != nil {
			panic(err)
		}

		handler, err := http.NewHandler(appService)
		if err != nil {
			panic(err)
		}
		server := &nethttp.Server{
			Handler: handler,
			Addr:    ":8080",
		}

		g.Add(func() error {
			httpLogger.Log("msg", "Start listening on port 8080")
			if err := server.ListenAndServe(); err != nethttp.ErrServerClosed {
				httpLogger.Log("msg", "Fail on ListenAndServe", "err", err)
				return err
			}

			return nil
		}, func(error) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				httpLogger.Log("msg", "Fail to shutdown", "err", err)
			}
		})

	}

	if err := g.Run(); err != nil {
		mainLogger.Log("msg", "Fail on run", "err", err)
	}

	mainLogger.Log("msg", "Stoppped")
}
