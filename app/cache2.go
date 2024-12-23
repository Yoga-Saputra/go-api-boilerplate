package app

import (
	"context"
	"log"

	rcache "github.com/go-redis/cache/v8"
	goredis "github.com/go-redis/redis/v8"

	"github.com/Yoga-Saputra/go-boilerplate/config"
	"github.com/Yoga-Saputra/go-boilerplate/pkg/redis"
)

// Cache2 driver pointer value
var Cache2 *cache2Adapter

type cache2Adapter struct {
	RCln *goredis.Client
	RCch *rcache.Cache
}

// Start cache redis connection
func cache2Up(args *AppArgs) {
	rda := redis.New(redis.Config{
		Host:       config.Of.Cache.Redis.Host,
		Port:       config.Of.Cache.Redis.Port,
		Password:   config.Of.Cache.Redis.Password,
		Database:   config.Of.Cache.Redis.Database,
		MaxRetries: config.Of.Cache.Redis.MaxRetries,
		PoolSize:   config.Of.Cache.Redis.PoolSize,
	})

	// Create new redis-cache instance
	cache2 := rcache.New(&rcache.Options{
		Redis: rda.Client,
	})

	// Create adapter
	Cache2 = &cache2Adapter{
		RCln: rda.Client,
		RCch: cache2,
	}
	printOutUp("New Cache2 Redis connection successfully open")
}

// Stop cache redis connection
func cache2Down() {
	printOutDown("Closing current Cache2 Redis connection...")

	if Cache2.RCln != nil {
		id := Cache2.RCln.ClientID(context.Background())

		if err := Cache2.RCln.Close(); err != nil {
			log.Printf("ERROR - failed to close redis connection, err: %v \n", err.Error())
		}

		log.Printf("SUCCESS - Redis connection already closed, %v \n", id)
	}
}
