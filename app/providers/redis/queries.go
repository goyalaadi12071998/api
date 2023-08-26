package redisclient

import (
	"context"
	"time"
)

type IRedisClient interface {
	Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error
}

func (r RedisClient) Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	err := r.client.Set(ctx, key, val, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}
