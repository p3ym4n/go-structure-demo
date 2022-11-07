package v1

import (
	"context"
	"go-structure-demo/internal/config"
	"go-structure-demo/internal/controller"
	"go-structure-demo/internal/log"
	"go-structure-demo/internal/param"
	"go-structure-demo/internal/pubsub"
	"go-structure-demo/internal/repository/postgresrepo"
	"go-structure-demo/internal/repository/redisrepo"
	"go-structure-demo/internal/validator"
)

func CreateUser(cfg *config.Config, logger log.Logger, redisRepo *redisrepo.RedisRepo, postgresRepo *postgresrepo.PostgresRepo) pubsub.MessageHandler {
	return func(ctx context.Context, bytes []byte) (bool, error) {
		event := new(UserCreatedEvent)

		// un marshall the event
		err := unmarshalPubSubEvent(bytes, event)
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

type UserCreatedEvent struct {
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

func unmarshalPubSubEvent([]byte, *UserCreatedEvent) error {
	return nil
}
