package cache

import (
	"context"
	"time"

	redis "github.com/go-redis/redis/v8"
)

type Client interface {
	Pipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error)
	HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	HGetAll(ctx context.Context, key string) *redis.StringStringMapCmd
	HMGet(ctx context.Context, key string, fields ...string) *redis.SliceCmd
	Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd
}

type Pipeliner = redis.Pipeliner

func NewClient(host string, password string, db int) (Client, error) {

	newClient := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
	})

	return newClient, nil
}
