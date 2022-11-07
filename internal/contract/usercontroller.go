package contract

import (
	"context"
	"go-structure-demo/internal/param"
)

type UserController interface {
	CreateUser(ctx context.Context, requestParam *param.CreateUserRequest) param.CreateUserResponse
}
