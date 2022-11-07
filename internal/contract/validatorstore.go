package contract

import (
	"context"
)

type ValidatorStore interface {
	Unique(ctx context.Context, entity, field string, value interface{}) (bool, error)
	Exists(ctx context.Context, entity, field string, value interface{}) (bool, error)
}
