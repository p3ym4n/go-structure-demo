package contract

import (
	"context"
	"go-structure-demo/internal/entity"
	"go-structure-demo/internal/param"
)

type UserStore interface {
	CreateUser(ctx context.Context, createUserRequest *param.CreateUserRequest) (entity.User, error)
}
