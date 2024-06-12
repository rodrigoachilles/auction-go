package user_entity

import (
	"context"
	"github.com/google/uuid"
	"github.com/rodrigoachilles/auction-go/internal/internal_error"
)

type User struct {
	Id   string
	Name string
}

type UserRepositoryInterface interface {
	CreateUser(
		ctx context.Context,
		user *User) *internal_error.InternalError

	FindUserById(
		ctx context.Context,
		userId string) (*User, *internal_error.InternalError)

	FindUsers(
		ctx context.Context) ([]User, *internal_error.InternalError)
}

func CreateUser(name string) (*User, *internal_error.InternalError) {
	user := &User{
		Id:   uuid.New().String(),
		Name: name,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Validate() *internal_error.InternalError {
	if len(u.Name) <= 1 {
		return internal_error.NewBadRequestError("invalid user object")
	}

	return nil
}
