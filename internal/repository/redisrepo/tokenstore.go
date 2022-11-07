package redisrepo

import (
	"context"
	"time"
)

func (rr *RedisRepo) SetToken(ctx context.Context, token string, userID uint, expiration time.Duration) error {
	return nil
}

func (rr *RedisRepo) PullToken(ctx context.Context, token string) (uint, bool) {
	return uint(0), true
}
