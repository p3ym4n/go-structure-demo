package param

import (
	v1 "go-structure-demo/internal/delivery/pubsub/handler/v1"
	"go-structure-demo/internal/entity"
	"net/http"
)

type CreateUserRequest struct {
	ID        uint    `form:"id" json:"id"`
	Email     string  `form:"email" json:"email"`
	FirstName string  `form:"first_name" json:"first_name"`
	LastName  string  `form:"last_name" json:"last_name"`
	Gender    *string `form:"gender" json:"gender"`
}

func (r *CreateUserRequest) BindFromChi(ctx *http.Request) error {
	return nil
}

func (r *CreateUserRequest) BindFromPubSub(event *v1.UserCreatedEvent) error {
	return nil
}

type CreateUserResponse struct {
	Message    string      `json:"message"`
	User       entity.User `json:"user"`
	Error      error       `json:"-"`
	StatusCode int         `json:"-"`
}

func (r *CreateUserResponse) ToJson() []byte {
	return []byte("do some marshalling")
}
