package subscriber

import (
	"context"
	"github.com/goflink/rider-workforce-common/log"
	"github.com/goflink/rider-workforce-common/pubsub"
	"go-structure-demo/internal/config"
	v1 "go-structure-demo/internal/delivery/pubsub/handler/v1"
	"go-structure-demo/internal/repository/postgresrepo"
	"go-structure-demo/internal/repository/redisrepo"
)

func Subscribe(ctx context.Context, cfg *config.Config, logger log.Logger, redisRepo *redisrepo.RedisRepo, postgresRepo *postgresrepo.PostgresRepo, pubsubClientA *pubsub.GCPClient, pubsubClientB *pubsub.GCPClient) func() {

	//create the clients
	pubsubClientA.Consume(ctx, cfg.PubSub.EmployeeHiredSubscriptionID, v1.CreateUser(cfg, logger, redisRepo, postgresRepo))

	return func() {
		//close pubsub client
		//pubsubClientA.Close()
		//pubsubClientB.Close()
	}
}
