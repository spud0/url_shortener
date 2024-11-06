package cache  

import (
	"context"
	"github.com/redis/v9"
	"net/http"
	"os"
)

// Caching scaffolding
type Cache struct {
	Client *redis.Client
}

func NewCache () *Cache {
	addr, passwd := os.Getenv("ADDR"), os.Getenv("CPASS")
	db := os.Getenv("DB_NUM")

	redisClient := redis.NewClient(&redis.Options {
		Addr: addr,
		Password: passwd,
		DB: db,
	})
	
	return &Cache{Client: redisClient}	

}

// Open a connection to the metrics platform to decide if the URL should be cached
func (c *Cache) CacheURL (ctx context.Context, hashedKey, longURL string) error {

	if ... {
		return c.Client.Set(
	} 

	return ...
}

// Returns a valid string or something else based on whether something is in the cache
func (c *Cache) GetURLFromCache (ctx context.Context, hashedKey string) (string, error) {
	return c.Client.Get(ctx, hashedKey).Result()
}
