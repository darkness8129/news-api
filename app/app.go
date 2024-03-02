package app

import (
	httpcontroller "darkness8129/news-api/app/controller/http"
	"darkness8129/news-api/app/entity"
	"darkness8129/news-api/app/service"
	"darkness8129/news-api/app/storage"
	"darkness8129/news-api/config"
	"darkness8129/news-api/packages/database"
	"darkness8129/news-api/packages/httpserver"
	"darkness8129/news-api/packages/logging"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Start(cfg *config.Config, logger logging.Logger) {
	logger = logger.Named("app")

	// connect to DB
	sql, err := database.NewPostgreSQLDatabase(database.Options{
		User:     cfg.PostgreSQL.User,
		Password: cfg.PostgreSQL.Password,
		Database: cfg.PostgreSQL.Database,
		Host:     cfg.PostgreSQL.Host,
		Logger:   logger,
	})
	if err != nil {
		logger.Fatal("failed to init postgresql db", "err", err)
	}

	db, ok := sql.DB().(*gorm.DB)
	if !ok {
		logger.Fatal("failed type assertion for db")
	}

	err = db.AutoMigrate(&entity.Post{})
	if err != nil {
		logger.Fatal("automigration failed", "err", err)
	}

	// init storages and services
	storages := service.Storages{
		Post: storage.NewPostStorage(db, logger),
	}
	services := service.Services{
		Post: service.NewPostService(storages, logger),
	}

	// init http server and start it
	httpServer := httpserver.NewGinHTTPServer(httpserver.Options{
		Addr:         cfg.HTTP.Addr,
		WriteTimeout: cfg.HTTP.WriteTimeout,
		ReadTimeout:  cfg.HTTP.ReadTimeout,
	})

	router := httpServer.Router().(*gin.Engine)
	if !ok {
		logger.Fatal("failed type assertion for router")
	}

	httpcontroller.New(httpcontroller.Options{
		Router:   router,
		Services: services,
		Logger:   logger,
	})

	httpServer.Start()

	// graceful shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info("app interrupt", "signal", s.String())
	case err := <-httpServer.Notify():
		logger.Error("err from notify ch", "err", err)
	}

	err = httpServer.Shutdown(cfg.ShutdownTimeout)
	if err != nil {
		logger.Error("failed to shutdown server", "err", err)
	}

	err = sql.Close()
	if err != nil {
		logger.Error("failed to close db connection", "err", err)
	}

	logger.Info("successful shutdown")
}
