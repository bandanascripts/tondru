package redis

import (
	"context"
	"fmt"
	"time"
)

func SetToRedis(ctx context.Context, key, value string, ttls int) error {

	if err := GlobalRedisClient.Set(ctx, key, value, time.Duration(ttls)*time.Second).Err(); err != nil {
		return fmt.Errorf("failed to set data to redis : %w", err)
	}

	return nil
}

func GetFromRedis(ctx context.Context, key string) (string, error) {

	result, err := GlobalRedisClient.Get(ctx, key).Result()

	if err != nil {
		return "", fmt.Errorf("failed to get data from redis : %w", err)
	}

	return result, nil
}

func PushToRedis(ctx context.Context, key, value string) error {

	if err := GlobalRedisClient.LPush(ctx, key, value).Err(); err != nil {
		return fmt.Errorf("failed to set data to redis : %w", err)
	}

	return nil
}

func RangeFromRedis(ctx context.Context, key string, start, end int64) ([]string, error) {

	result, err := GlobalRedisClient.LRange(ctx, key, start, end).Result()

	if err != nil {
		return nil, fmt.Errorf("failed to fetch results from redis : %w", err)
	}

	return result, nil
}
