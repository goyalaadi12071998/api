package redisclient

import (
	"context"
	"time"
)

type IRedisQuery interface {
	Set()
}

func (r RedisClient) Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	err := r.client.Set(ctx, key, val, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r RedisClient) Get(ctx context.Context, key string) (string, error) {
	data := r.client.Get(ctx, key)
	if data.Err() != nil {
		return "", data.Err()
	}
	value := data.Val()
	return value, nil
}
