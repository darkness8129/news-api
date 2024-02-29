package httpserver

type HTTPServer interface {
	Start()
	Notify() <-chan error
	Router() interface{}
	Shutdown() error
}
