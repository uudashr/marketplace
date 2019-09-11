package main

import (
	"context"
	"flag"
	"fmt"
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
	flagListenAddr := flag.String("http-addr", ":8080", "HTTP listen address")
	flagLogFormat := flag.String("log-format", "logfmt", "Log format (logfmt|json)")
	flagHelp := flag.Bool("help", false, "Show help")
	flag.Parse()

	if *flagHelp {
		flag.Usage()
		os.Exit(0)
	}

	logWriter := log.NewSyncWriter(os.Stderr)
	logger, err := logWithFormat(logWriter, *flagLogFormat)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		flag.Usage()
		os.Exit(1)
	}

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
		storeRepo := inmem.NewStoreRepository()
		appService, err := app.NewService(categoryRepo, storeRepo)
		if err != nil {
			panic(err)
		}

		handler, err := http.NewHandler(appService)
		if err != nil {
			panic(err)
		}

		server := &nethttp.Server{
			Handler: handler,
			Addr:    *flagListenAddr,
		}

		g.Add(func() error {
			httpLogger.Log("msg", "Start listening", "addr", *flagListenAddr)
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
