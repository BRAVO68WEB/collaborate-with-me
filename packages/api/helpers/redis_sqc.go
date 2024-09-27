package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/vektah/gqlparser/v2/ast"
)

type CacheSQC struct {
	client redis.UniversalClient
	ttl    time.Duration
}

const sqcPrefix = "sqc:"

func NewSQCCache(redisAddress string, ttl time.Duration) (*CacheSQC, error) {
	client := redis.NewClient(&redis.Options{
		Addr: redisAddress,
	})

	err := client.Ping().Err()
	if err != nil {
		return nil, fmt.Errorf("could not create cache: %w", err)
	}

	return &CacheSQC{client: client, ttl: ttl}, nil
}

func (c *CacheSQC) Add(ctx context.Context, key string, value *ast.QueryDocument) {
	c.client.Set(apqPrefix+key, value, c.ttl)
}

func (c *CacheSQC) Get(ctx context.Context, key string) (*ast.QueryDocument, bool) {
	s, err := c.client.Get(apqPrefix + key).Result()
	if err != nil {
		return nil, false
	}
	var queryDoc ast.QueryDocument
	err = json.Unmarshal([]byte(s), &queryDoc)
	if err != nil {
		return nil, false
	}
	return &queryDoc, true
}
