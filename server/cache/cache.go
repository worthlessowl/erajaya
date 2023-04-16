package cache

import (
	"time"

	"github.com/go-redis/redis"
)

type cacheInterface interface {
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(key string) *redis.StringCmd
}

type Cache struct {
	client cacheInterface
}

func New(client cacheInterface) Cache {
	return Cache{client: client}
}
