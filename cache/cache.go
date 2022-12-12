package cache

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Cache struct {
	redis *redis.Client
}

func (cache *Cache) Add(consumer, harvester string, bandwidth int64) {
	cache.redis.HSet(ctx, consumer, harvester, bandwidth)
}

func (cache *Cache) All() error {
	iter := cache.redis.Scan(ctx, 0, "*", 0).Iterator()
	var err error
	for iter.Next(ctx) {
		consumer := iter.Val()
		harvesters, err := cache.redis.HGetAll(ctx, consumer).Result()
		if err != nil {
			return err
		}
		for harvester, bandwidth_str := range harvesters {
			bandwidth, err := strconv.ParseInt(bandwidth_str, 10, 64)
			if err != nil {
				return err
			}
			fmt.Println(consumer, harvester, bandwidth)
		}
	}
	if err = iter.Err(); err != nil {
		return err
	}
	return nil
}

func NewCache(addr, username, password string, db int) *Cache {
	redis := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: username,
		Password: password,
		DB:       db,
	})

	cache := Cache{
		redis: redis,
	}

	return &cache
}
