package v1

import (
	"context"
	"github.com/goflink/events/go/hrpb"
	"github.com/goflink/rider-workforce-common/log"
	"github.com/goflink/rider-workforce-common/pubsub"
	"github.com/golang/protobuf/proto"
	"go-structure-demo/internal/config"
	"go-structure-demo/internal/controller"
	"go-structure-demo/internal/param"
	"go-structure-demo/internal/repository/postgresrepo"
	"go-structure-demo/internal/repository/redisrepo"
	"go-structure-demo/internal/validator"
)

func CreateUser(cfg *config.Config, logger log.Logger, redisRepo *redisrepo.RedisRepo, postgresRepo *postgresrepo.PostgresRepo) pubsub.MessageHandler {
	return func(ctx context.Context, bytes []byte) (bool, error) {
		event := new(hrpb.Auth0IdentityCreated)
		err := proto.Unmarshal(bytes, event)
		if err != nil {
			return true, nil
		}

		requestDTO := new(param.CreateUserRequest)
		err = requestDTO.BindFromPubSub(event)
		if err != nil {
			return false, err
		}

		err = validator.CreateUserRequest(ctx, requestDTO, postgresRepo)
		if err != nil {
			return false, err
		}

		userCreateResponse := controller.NewUserController(postgresRepo).CreateUser(ctx, requestDTO)
		if userCreateResponse.Error != nil {
			return false, userCreateResponse.Error
		}

		return true, nil
	}
}
