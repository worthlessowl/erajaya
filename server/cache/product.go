package cache

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis"
)

func (cache *Cache) SetProductCache(key string, products []Product) error {
	marshalled, err := json.Marshal(products)
	if err != nil {
		log.Println("Failed to marshal product")
		return err
	}

	cacheErr := cache.client.Set(key, marshalled, 10*time.Second)
	if err != nil {
		log.Println("Failed to set cache")
		return cacheErr.Err()
	}
	return nil
}

func (cache *Cache) GetProductCache(key string) ([]Product, error) {
	value, err := cache.client.Get(key).Result()
	if err == redis.Nil {
		log.Println("Key not found")
		return nil, nil
	} else if err != nil {
		log.Println("Failed to get cache")
		return nil, err
	}

	products := []Product{}

	json.Unmarshal([]byte(value), &products)

	return products, nil
}
