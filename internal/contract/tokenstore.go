package contract

import (
	"context"
	"time"
)

type TokenStore interface {
	SetToken(ctx context.Context, token string, userID uint, expiration time.Duration) error
	PullToken(ctx context.Context, token string) (uint, bool)
}
