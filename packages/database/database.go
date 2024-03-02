package database

type Database interface {
	DB() interface{}
	Close() error
}
