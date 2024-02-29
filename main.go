package main

import (
	"darkness8129/news-api/app"
	"darkness8129/news-api/config"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	var cfg config.Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		fmt.Println("failed to read cfg")
		os.Exit(1)
	}

	app.Start(&cfg)
}
