package httpserver

import (
	"context"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

var _ HTTPServer = (*ginHTTPServer)(nil)

type ginHTTPServer struct {
	server   *http.Server
	router   *gin.Engine
	notifyCh chan error
}

type Options struct {
	Addr         string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}

func NewGinHTTPServer(opt Options) *ginHTTPServer {
	router := gin.New()

	httpServer := &http.Server{
		Handler:      router,
		Addr:         opt.Addr,
		WriteTimeout: opt.WriteTimeout,
		ReadTimeout:  opt.ReadTimeout,
	}

	return &ginHTTPServer{
		server:   httpServer,
		router:   router,
		notifyCh: make(chan error, 1),
	}
}

func (s *ginHTTPServer) Start() {
	go func() {
		defer close(s.notifyCh)
		s.notifyCh <- s.server.ListenAndServe()
	}()
}

func (s *ginHTTPServer) Notify() <-chan error {
	return s.notifyCh
}

func (s *ginHTTPServer) Router() interface{} {
	return s.router
}

func (s *ginHTTPServer) Shutdown(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
