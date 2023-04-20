package cache

import (
	"time"
)

type Stub struct {
	store map[string]string
}

func InitStub() {
	Instance = &Stub{store: map[string]string{}}
}

func (c Stub) SetValue(key, value string, expiration time.Duration) bool {
	c.store[key] = value
	return true
}

func (c Stub) DeleteValue(key string) {
	delete(c.store, key)
}

func (c Stub) FetchValue(key string) string {
	return c.store[key]
}

var _ Client = &Stub{}
