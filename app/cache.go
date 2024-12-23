package app

/* import (
	"github.com/Yoga-Saputra/go-boilerplate/config"
	"github.com/Yoga-Saputra/go-boilerplate/pkg/redis"
)

// Cache driver pointer value for redis adapter
var Cache *redis.Redis

// Start cache redis connection
func cacheUp(args *AppArgs) {
	redis := redis.New(redis.Config{
		Host:       config.Of.Cache.Redis.Host,
		Port:       config.Of.Cache.Redis.Port,
		Password:   config.Of.Cache.Redis.Password,
		Database:   config.Of.Cache.Redis.Database,
		MaxRetries: config.Of.Cache.Redis.MaxRetries,
		PoolSize:   config.Of.Cache.Redis.PoolSize,
		NameSpace:  "republish-seamless-wallet:cache",
	})

	Cache = redis
	printOutUp("New Cache Redis connection successfully open")
}

// Stop cache redis connection
func cacheDown() {
	printOutDown("Closing current Cache Redis connection...")
	Cache.Close(false)
} */
