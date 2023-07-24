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

func (p RedisClient) SetKeyWithExpiry(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	p.pipeline.SAdd(ctx, key, val)
	p.pipeline.Expire(ctx, key, expiration)
	_, err := p.pipeline.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p RedisClient) GetToalRequestForUser(ctx context.Context, key string) (int, error) {
	totalRequest, err := p.client.SCard(ctx, key).Result()
	if err != nil {
		return -1, err
	}

	return int(totalRequest), nil
}
