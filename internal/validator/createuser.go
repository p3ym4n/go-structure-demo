package validator

import (
	"context"
	"go-structure-demo/internal/contract"
	"go-structure-demo/internal/param"
)

func CreateUserRequest(ctx context.Context, dto *param.CreateUserRequest, store contract.ValidatorStore) error {
	// doing the validations and return error
	return nil
}
