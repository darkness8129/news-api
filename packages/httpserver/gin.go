package httpserver

import (
	"context"
	"darkness8129/news-api/config"

	"net/http"

	"github.com/gin-gonic/gin"
)

var _ HTTPServer = (*ginHTTPServer)(nil)

type ginHTTPServer struct {
	cfg    *config.Config
	server *http.Server
	notify chan error
}

func NewGinHTTPServer(cfg *config.Config) *ginHTTPServer {
	router := gin.New()

	httpServer := &http.Server{
		Handler:      router,
		Addr:         cfg.HTTP.Addr,
		WriteTimeout: cfg.HTTP.WriteTimeout,
		ReadTimeout:  cfg.HTTP.ReadTimeout,
	}

	return &ginHTTPServer{
		cfg:    cfg,
		server: httpServer,
		notify: make(chan error, 1),
	}
}

func (s *ginHTTPServer) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *ginHTTPServer) Notify() <-chan error {
	return s.notify
}

func (s *ginHTTPServer) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.cfg.ShutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
