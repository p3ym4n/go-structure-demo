package redisrepo

import "go-structure-demo/internal/config"

type RedisRepo struct {
}

func New(cfg *config.Config) (*RedisRepo, func()) {
	return &RedisRepo{}, func() {}
}
