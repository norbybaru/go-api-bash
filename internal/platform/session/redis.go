package session

import (
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisClient struct {
	conn *redis.Client
}

// initialize redis session
func initRedisStore(host string, port string, password string, databaseCount int) Storage {
	client, err := redisConnection(host, port, password, databaseCount)

	if err != nil {
		log.Fatal("ERR - Redis conn: ", err)
	}

	return NewRedisStorage(client)
}

// RedisConnection func for connect to Redis server.
func redisConnection(host string, port string, password string, databaseCount int) (*redis.Client, error) {
	// URL for Redis connection.
	connURL := fmt.Sprintf(
		"%s:%s",
		host,
		port,
	)

	// Set Redis options.
	options := &redis.Options{
		Addr:     connURL,
		Password: password,
		DB:       databaseCount,
	}

	return redis.NewClient(options), nil
}

func NewRedisStorage(redis *redis.Client) Storage {
	return &redisClient{redis}
}

// Store key, value pair in redis
func (r *redisClient) Set(key string, value string, ttl time.Duration) error {
	return r.conn.Set(ctx, key, value, ttl).Err()
}

// Retrieve stored key, return default if key missing
func (r *redisClient) Get(key string, defaultValue string) (interface{}, error) {
	return r.conn.Get(ctx, key).Result()
}

func (r *redisClient) Delete(key string) error {
	return r.conn.Del(ctx, key).Err()
}

func (r *redisClient) Close() error {
	return r.conn.Close()
}
