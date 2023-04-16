package cache

import (
	"fmt"

	"github.com/go-redis/redis"
)

const (
	host = "redis"
	port = 6379
)

func InitRedis() *redis.Client {
	redisHost := fmt.Sprintf("%s:%d", host, port)

	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "",
		DB:       0,
	})

	return client
}
