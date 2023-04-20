package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var ctx = context.Background()

type Client interface {
	FetchValue(key string) string
	SetValue(key, value string) bool
}

type RedisClient struct {
	instance *redis.Client
}

func InitClient(host, port, password string) Client {
	i := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0, // use default DB
	})
	return &RedisClient{i}
}

func (c *RedisClient) FetchValue(key string) string {
	val, err := c.instance.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return val
}

func (c *RedisClient) SetValue(key, value string) bool {
	err := c.instance.Set(ctx, key, value, 24*time.Hour).Err()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
