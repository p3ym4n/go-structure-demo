package controller

import (
	"context"
	"go-structure-demo/internal/contract"
	"go-structure-demo/internal/param"
	"net/http"
)

type UserController struct {
	userStore contract.UserStore
}

func NewUserController(userStore contract.UserStore) *UserController {
	return &UserController{
		userStore: userStore,
	}
}

func (c *UserController) CreateUser(ctx context.Context, request *param.CreateUserRequest) param.CreateUserResponse {
	//call everything needed, validation, authorization, creation,...
	return param.CreateUserResponse{
		Message:    "user created!",
		Error:      nil,
		StatusCode: http.StatusCreated,
	}
}
