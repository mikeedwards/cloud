package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/kelseyhightower/envconfig"

	_ "github.com/lib/pq"

	"github.com/fieldkit/cloud/server/ingester"
	"github.com/fieldkit/cloud/server/logging"
	"github.com/fieldkit/cloud/server/metrics"
)

func main() {
	ctx := context.Background()

	config := getConfig()

	logging.Configure(config.ProductionLogging, "ingester")

	log := logging.Logger(ctx).Sugar()

	log.Info("starting")

	newIngester := ingester.NewIngester(ctx, config)

	notFoundHandler := http.NotFoundHandler()

	coreHandler := configureMetrics(ctx, config, http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/status" {
			log.Infow("status", "headers", req.Header)
			fmt.Fprint(w, "ok")
			return
		}

		if req.URL.Path == "/ingestion" {
			newIngester.ServeHTTP(w, req)
			return
		}

		notFoundHandler.ServeHTTP(w, req)
	}))

	server := &http.Server{
		Addr:    config.Addr,
		Handler: coreHandler,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Errorw("startup", "err", err)
	}
}

// I'd like to make this common with server where possible.

func configureMetrics(ctx context.Context, config *ingester.Config, next http.Handler) http.Handler {
	log := logging.Logger(ctx).Sugar()

	if config.StatsdAddress == "" {
		log.Infow("stats: skipping")
		return next
	}

	log.Infow("stats", "address", config.StatsdAddress)

	metricsSettings := metrics.MetricsSettings{
		Prefix:  "fk.ingester",
		Address: config.StatsdAddress,
	}

	return metrics.GatherMetrics(ctx, metricsSettings, next)
}

func getConfig() *ingester.Config {
	var config ingester.Config

	flag.BoolVar(&config.Help, "help", false, "usage")

	flag.Parse()

	if config.Help {
		flag.Usage()
		envconfig.Usage("fieldkit", &config)
		os.Exit(0)
	}

	err := envconfig.Process("fieldkit", &config)
	if err != nil {
		panic(err)
	}

	return &config
}
