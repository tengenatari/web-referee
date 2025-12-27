package redisstorage

import (
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisStorage struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisStorage(addr string) *RedisStorage {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		PoolSize: 100,
	})

	redisStorage := &RedisStorage{
		client: client,
		ttl:    24 * time.Hour,
	}

	return redisStorage
}
