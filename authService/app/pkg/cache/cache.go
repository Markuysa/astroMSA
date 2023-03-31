package cache

import (
	"context"
	"errors"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	CacheElemAddingErr = errors.New("cache adding err")
)

type Cache struct {
	Client *cache.Cache
}

func New(opts *redis.Client) *Cache {
	return &Cache{
		Client: cache.New(&cache.Options{Redis: opts}),
	}
}

func (c *Cache) Add(ctx context.Context, obj interface{}, key string) error {
	err := c.Client.Set(&cache.Item{
		Key:   key,
		Ctx:   ctx,
		Value: obj,
		TTL:   time.Duration(time.Hour * 24),
	})
	if err != nil {
		return CacheElemAddingErr
	}
	return nil
}

func (c *Cache) Get(ctx context.Context, key string) (interface{}, error) {

	var res interface{}
	err := c.Client.Get(ctx, key, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
