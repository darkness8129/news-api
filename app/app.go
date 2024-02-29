package app

import (
	"darkness8129/news-api/config"
	"darkness8129/news-api/packages/httpserver"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Start(cfg *config.Config) {
	httpServer := httpserver.NewGinHTTPServer(cfg)
	httpServer.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		fmt.Println("app interrupt: ", s)
	case err := <-httpServer.Notify():
		fmt.Println("err from notify ch: ", err)
	}

	err := httpServer.Shutdown()
	if err != nil {
		fmt.Println("failed to shutdown")
	}
}
