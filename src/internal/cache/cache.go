package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var ctx = context.Background()

type Client interface {
	FetchValue(key string) string
	SetValue(key, value string, expiration time.Duration) bool
	DeleteValue(key string)
}

type RedisClient struct {
	instance *redis.Client
}

var Instance Client

func InitClient(host, port string) {
	i := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
		DB:   0, // use default DB,
	})
	Instance = &RedisClient{i}
}

func (c *RedisClient) FetchValue(key string) string {
	val, err := c.instance.Get(ctx, key).Result()
	if err != nil {
		log.Printf("Error during fetching value %s: %s\n", key, err)
		return ""
	}
	return val
}

func (c *RedisClient) SetValue(key, value string, expiration time.Duration) bool {
	err := c.instance.Set(ctx, key, value, expiration).Err()
	if err != nil {
		log.Printf("Error during setting value %s: %s\n", key, err)
		return false
	}
	return true
}

func (c *RedisClient) DeleteValue(key string) {
	err := c.instance.Del(ctx, key).Err()
	if err != nil {
		log.Printf("Error during deleting value %s: %s\n", key, err)
	}
}

var _ Client = &RedisClient{}
