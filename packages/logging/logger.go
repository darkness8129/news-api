package logging

type Logger interface {
	Named(name string) Logger
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Error(message string, args ...interface{})
	Fatal(message string, args ...interface{})
}
