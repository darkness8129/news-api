package main

import (
	"darkness8129/news-api/app"
	"darkness8129/news-api/config"
	"darkness8129/news-api/packages/logging"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	logger, err := logging.NewZapLogger()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var cfg config.Config
	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		logger.Fatal("failed to read config", "err", err)
	}

	app.Start(&cfg, logger)
}
