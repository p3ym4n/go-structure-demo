package param

import (
	"github.com/goflink/events/go/hrpb"
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

func (r *CreateUserRequest) BindFromPubSub(event *hrpb.Auth0IdentityCreated) error {
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
