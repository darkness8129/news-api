package main

import (
	"darkness8129/news-api/app"
	"darkness8129/news-api/config"
	"darkness8129/news-api/packages/logging"
	"log"
	"os"
)

func main() {
	logger, err := logging.NewZapLogger()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	cfg, err := config.New()
	if err != nil {
		logger.Fatal("failed to get config", "err", err)
	}

	app.Start(cfg, logger)
}
