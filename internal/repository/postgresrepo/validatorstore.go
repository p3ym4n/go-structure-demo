package postgresrepo

import (
	"context"
)

func (p *PostgresRepo) Unique(ctx context.Context, entity, field string, value interface{}) (bool, error) {
	return false, nil
}

func (p *PostgresRepo) Exists(ctx context.Context, entity, field string, value interface{}) (bool, error) {
	return false, nil
}
