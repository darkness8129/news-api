package app

import (
	httpcontroller "darkness8129/news-api/app/controller/http"
	"darkness8129/news-api/app/entity"
	"darkness8129/news-api/app/service"
	"darkness8129/news-api/app/storage"
	"darkness8129/news-api/config"
	"darkness8129/news-api/packages/database"
	"darkness8129/news-api/packages/httpserver"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func Start(cfg *config.Config) {
	db, err := database.NewPostgreSQLDatabase(cfg)
	if err != nil {
		os.Exit(1)
		fmt.Println("failed to init postgresql db: ", err)
	}

	err = db.DB.AutoMigrate(&entity.Post{})
	if err != nil {
		os.Exit(1)
		fmt.Println("automigration failed: ", err)
	}

	storages := service.Storages{
		Post: storage.NewPostService(db),
	}

	services := service.Services{
		Post: service.NewPostService(storages),
	}

	httpServer := httpserver.NewGinHTTPServer(cfg)
	router := httpServer.Router().(*gin.Engine)

	httpcontroller.New(httpcontroller.Options{
		Router:   router,
		Services: services,
	})

	httpServer.Start()

	// graceful shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		fmt.Println("app interrupt: ", s)
	case err := <-httpServer.Notify():
		fmt.Println("err from notify ch: ", err)
	}

	err = httpServer.Shutdown()
	if err != nil {
		fmt.Println("failed to shutdown server: ", err)
	}

	err = db.Close()
	if err != nil {
		fmt.Println("failed to close db connection: ", err)
	}
}
