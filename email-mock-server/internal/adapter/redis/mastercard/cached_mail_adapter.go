package redis

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/redbeestudios/go-parser-poc/commons/pkg"
	"log"
	"time"
)

const MailKey = "mail"

type CachedMailAdapter struct {
	cache *redis.Client
}

func (a *CachedMailAdapter) Save(mail *Mail) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 6*time.Second)
	defer cancel()

	status := a.cache.Set(ctx, keyForName(mail.To), mail.Body, 0)

	if status.Err() != nil {
		log.Println(status.Err())
	}

	return
}

func NewCachedMailAdapter(
	cacheConfig *pkg.EmailMockServerConfig,
) *CachedMailAdapter {
	cache := redis.NewClient(&redis.Options{
		Addr:     cacheConfig.Redis.Address,
		Password: "",
		DB:       0,
	})

	return &CachedMailAdapter{
		cache: cache,
	}
}

func keyForName(name string) string {
	return MailKey + ":" + name
}
