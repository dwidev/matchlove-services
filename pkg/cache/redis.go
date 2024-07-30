package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewRedisClient() *Redis {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	redisClient.Ping(context.Background())

	return &Redis{redisClient}
}

type Redis struct {
	client *redis.Client
}

func (r *Redis) Disconnect(ctx context.Context) error {
	err := r.client.Close()
	if err != nil {
		return err
	}

	return nil
}

func (r *Redis) GetString(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (r *Redis) SetString(ctx context.Context, key string, value string, expiration time.Duration) error {
	err := r.client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Redis) Delete(ctx context.Context, key string) error {
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}
