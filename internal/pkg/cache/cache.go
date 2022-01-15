package cache

import (
	"os"
	"sync"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type Item = cache.Item

var (
	redisCacheOnce sync.Once
	redisCache     *cache.Cache
	Once           = RedisCache().Once
	Set            = RedisCache().Set
	Get            = RedisCache().Get
	Delete         = RedisCache().Delete
)

func RedisCache() *cache.Cache {
	redisCacheOnce.Do(func() {
		redisCache = cache.New(&cache.Options{
			Redis: redis.NewRing(&redis.RingOptions{
				Addrs: map[string]string{
					"redis": os.Getenv("REDIS_CACHE_SERVICE_HOST") + ":" + os.Getenv("REDIS_CACHE_SERVICE_PORT"),
				},
			}),
			LocalCache: cache.NewTinyLFU(1000, time.Minute),
		})
	})
	return redisCache
}
