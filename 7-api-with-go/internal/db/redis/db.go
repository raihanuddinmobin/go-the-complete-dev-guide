package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func Connect() (*redis.Client, error) {
	ctx := context.Background()

	opt, err := redis.ParseURL("redis://default:@localhost:6379/0")
	if err != nil {
		return nil, fmt.Errorf("failed to connect redis %w ", err)
	}

	client := redis.NewClient(opt)

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping redis %w", err)
	}

	fmt.Println("Successfully Connected Redis ✅")
	return client, err
}
