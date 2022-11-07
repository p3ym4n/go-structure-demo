package postgresrepo

import (
	"context"
	"go-structure-demo/internal/entity"
	"go-structure-demo/internal/param"
)

func (p *PostgresRepo) CreateUser(ctx context.Context, createUserRequest *param.CreateUserRequest) (entity.User, error) {
	return entity.User{}, nil
}
