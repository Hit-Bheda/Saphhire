package queue

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func Enqueue(ctx context.Context, rdb *redis.Client, url string) error {
	return rdb.RPush(ctx, "crawler:queue", url).Err()
}

func Dequeue(ctx context.Context, rdb *redis.Client) (string, error) {
	return rdb.RPop(ctx, "crawler:queue").Result()
}
