package helpers

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type CacheAPQ struct {
	client redis.UniversalClient
	ttl    time.Duration
}

const apqPrefix = "apq:"

func NewAPQCache(redisAddress string, ttl time.Duration) (*CacheAPQ, error) {
	client := redis.NewClient(&redis.Options{
		Addr: redisAddress,
	})

	err := client.Ping().Err()
	if err != nil {
		return nil, fmt.Errorf("could not create cache: %w", err)
	}

	return &CacheAPQ{client: client, ttl: ttl}, nil
}

func (c *CacheAPQ) Add(ctx context.Context, key string, value string) {
	c.client.Set(apqPrefix+key, value, c.ttl)
}

func (c *CacheAPQ) Get(ctx context.Context, key string) (string, bool) {
	s, err := c.client.Get(apqPrefix + key).Result()
	if err != nil {
		return "", false
	}
	return s, true
}
