package cache

import (
	"github.com/redis/go-redis/v9"
	"time"
)

type Stub struct {
	instance *redis.Client
}

func (c *Stub) FetchValue(key string) string {
	return ""
}

func (c *Stub) SetValue(key, value string, expiration time.Duration) bool {
	return false
}

func (c *Stub) DeleteValue(key string) {

}

var _ Client = &Stub{}
