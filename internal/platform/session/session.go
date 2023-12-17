package session

import (
	"context"
	"time"
)

type Storage interface {
	Set(key string, value string, ttl time.Duration) error
	Get(key string, defaultValue string) (interface{}, error)
	Delete(key string) error
	Close() error
}

type Session struct {
	store Storage
}

var ctx = context.Background()

func Init(host string, port string, password string) Storage {
	store := initRedisStore(host, port, password, 0)

	return store
}
