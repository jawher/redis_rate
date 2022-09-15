package redis_rate_test

import (
	"context"
	"fmt"

	"os"

	"github.com/go-redis/redis_rate/v10"

	"github.com/go-redis/redis/v9"
)

func ExampleNewLimiter() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})
	_ = rdb.FlushDB(ctx).Err()

	limiter := redis_rate.NewLimiter(rdb)
	res, err := limiter.Allow(ctx, "project:123", redis_rate.PerSecond(10))
	if err != nil {
		panic(err)
	}
	fmt.Println("allowed", res.Allowed, "remaining", res.Remaining)
	// Output: allowed 1 remaining 9
}
