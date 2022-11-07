package redis

import (
	"context"
	"time"

	goredis "github.com/go-redis/redis/v8"
)

type Adapter interface {
	Client() *goredis.Client
	Ping(ctx context.Context) (string, error)

	Set(ctx context.Context, key string, value interface{}, expiration time.Duration, tags ...string) error
	Forever(ctx context.Context, key string, value interface{}) error

	Has(ctx context.Context, key string) bool
	Get(ctx context.Context, key string) (string, bool)
	TTL(ctx context.Context, key string) (time.Duration, bool)

	Pull(ctx context.Context, key string) (interface{}, bool)
	Invalidate(ctx context.Context, tags ...string)
	Del(ctx context.Context, key string) bool
	Flush(ctx context.Context) error
}
