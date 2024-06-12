package user_usecase

import (
	"context"
	"github.com/rodrigoachilles/auction-go/internal/internal_error"
)

func (u *UserUseCase) FindUsers(
	ctx context.Context) ([]UserOutputDTO, *internal_error.InternalError) {
	userEntities, err := u.UserRepository.FindUsers(ctx)
	if err != nil {
		return nil, err
	}

	var userOutputs []UserOutputDTO
	for _, user := range userEntities {
		userOutputs = append(userOutputs, UserOutputDTO{
			Id:   user.Id,
			Name: user.Name,
		})
	}

	return userOutputs, nil
}

func (u *UserUseCase) FindUserById(
	ctx context.Context, id string) (*UserOutputDTO, *internal_error.InternalError) {
	userEntity, err := u.UserRepository.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &UserOutputDTO{
		Id:   userEntity.Id,
		Name: userEntity.Name,
	}, nil
}
