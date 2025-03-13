package database

type DB interface {
	Ping() error
	Close() error
	Set() error
}
