package private

import (
	"go-structure-demo/internal/pubsub"
	"go-structure-demo/internal/repository/postgresrepo"
	"go-structure-demo/internal/repository/redisrepo"
	"net/http"
)

func Health(
	redisRepo *redisrepo.RedisRepo,
	postgresRepo *postgresrepo.PostgresRepo,
	pubsubClientA *pubsub.GCPClient,
	pubsubClientB *pubsub.GCPClient,
) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}
