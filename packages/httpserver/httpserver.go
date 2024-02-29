package httpserver

type HTTPServer interface {
	Start()
	Notify() <-chan error
	Shutdown() error
}
