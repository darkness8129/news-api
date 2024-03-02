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
	"gorm.io/gorm"
)

func Start(cfg *config.Config) {
	// connect to DB
	sql, err := database.NewPostgreSQLDatabase(database.Options{
		User:     cfg.PostgreSQL.User,
		Password: cfg.PostgreSQL.Password,
		Database: cfg.PostgreSQL.Database,
		Host:     cfg.PostgreSQL.Host,
	})
	if err != nil {
		os.Exit(1)
		fmt.Println("failed to init postgresql db: ", err)
	}

	db, ok := sql.DB().(*gorm.DB)
	if !ok {
		os.Exit(1)
		fmt.Println("failed type assertion for db")
	}

	err = db.AutoMigrate(&entity.Post{})
	if err != nil {
		os.Exit(1)
		fmt.Println("automigration failed: ", err)
	}

	// init storages and services
	storages := service.Storages{
		Post: storage.NewPostService(db),
	}
	services := service.Services{
		Post: service.NewPostService(storages),
	}

	// init http server and start it
	httpServer := httpserver.NewGinHTTPServer(httpserver.Options{
		Addr:         cfg.HTTP.Addr,
		WriteTimeout: cfg.HTTP.WriteTimeout,
		ReadTimeout:  cfg.HTTP.ReadTimeout,
	})

	router := httpServer.Router().(*gin.Engine)
	if !ok {
		os.Exit(1)
		fmt.Println("failed type assertion for router")
	}

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

	err = httpServer.Shutdown(cfg.ShutdownTimeout)
	if err != nil {
		fmt.Println("failed to shutdown server: ", err)
	}

	err = sql.Close()
	if err != nil {
		fmt.Println("failed to close db connection: ", err)
	}
}
