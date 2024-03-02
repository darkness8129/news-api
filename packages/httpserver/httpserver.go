package httpserver

import "time"

type HTTPServer interface {
	Start()
	Notify() <-chan error
	Router() interface{}
	Shutdown(timeout time.Duration) error
}
