package cache

import (
	"context"
	"fmt"
	"time"
)

func New(t Type) Cache {
	if t == RedisCache {
		return NewRedisClient()
	}

	panic("not implement type for Run in cache")
}

type Cache interface {
	GetString(ctx context.Context, key string) (string, error)
	SetString(ctx context.Context, key string, value string, expiration time.Duration) error
	Delete(ctx context.Context, key string) error
	Disconnect(ctx context.Context) error
}

func AccessTokenKeyCache(accountID string, imei string) string {
	return fmt.Sprintf("access_token_%s_%s", accountID, imei)
}
