package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/nevzheng/hexlink"
	rr "github.com/nevzheng/hexlink/repository/redis"
	"github.com/nevzheng/hexlink/shortener"
)

func main() {
	// Configure Logger
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "hexlink",
			"time:", log.DefaultTimestampUTC,
			"caller:", log.DefaultCaller)
	}

	port := os.Getenv("PORT")
	if port == "" {
		level.Warn(logger).Log("msg", "Port Not Specified. Using default 8080")
		port = "8080"
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended") // Print on Exit

	// Instantiate context to use across calls
	ctx := context.Background()
	// Initialize the Service
	var srv shortener.RedirectService
	{
		// Set up Redis
		redisURL := os.Getenv("REDIS_URL")
		if redisURL == "" {
			defaultUrl := "redis://localhost:6379"
			level.Warn(logger).Log("msg", "REDIS_URL Not Specified. Using default", "url", defaultUrl)
			redisURL = defaultUrl
		}
		repo, err := rr.NewRedisRepository(redisURL, logger)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(1)
		}
		level.Info(logger).Log("msg", "Redis Client Created")

		// assign the service
		srv = shortener.NewRedirectService(repo, logger)
	}

	// Await and handle unrecoverable errors and conditions
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := hexlink.MakeEndpoints(srv)

	go func() {
		level.Info(logger).Log("msg", "Listening on Port: "+port)
		handler := hexlink.NewHttpServer(ctx, endpoints)
		errs <- http.ListenAndServe(":"+port, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
