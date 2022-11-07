package postgresrepo

import (
	"go-structure-demo/internal/config"
)

type PostgresRepo struct {
}

func New(cfg *config.Config) (*PostgresRepo, func()) {
	//open the connection
	return &PostgresRepo{}, func() {
		// pass the connections closer
	}
}
