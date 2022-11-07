package entity

import (
	"fmt"
	"strings"
)

const (
	UserEntity      = "users"
	UserEntityID    = "id"
	UserEntityEmail = "email"
)

const (
	GenderMale   = "male"
	GenderFemale = "female"
)

type User struct {
	ID        uint    `json:"id" uri:"user"`
	Email     string  `json:"email"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Gender    *string `json:"gender"`
}

func (user *User) GetFullName() string {
	return strings.TrimSpace(fmt.Sprintf("%s %s", user.FirstName, user.LastName))
}

func (user *User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":         user.ID,
		"email":      user.Email,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"gender":     user.Gender,
	}
}
