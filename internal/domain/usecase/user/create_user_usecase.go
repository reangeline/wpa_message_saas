package usecase

import (
	"context"
	"errors"

	repository "github.com/reangeline/micro_saas/internal/domain/contract/repository"
	"github.com/reangeline/micro_saas/internal/domain/entity"
	"github.com/reangeline/micro_saas/internal/dto"
)

type CreateUserUseCase struct {
	userRepository repository.UserRepositoryInterface
}

func NewCreateUserUseCase(
	userRepository repository.UserRepositoryInterface,
) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: userRepository,
	}
}

var (
	ErrEmailAlreadyExists = errors.New("email already exist")
)

func (u *CreateUserUseCase) Execute(ctx context.Context, input *dto.UserInput) error {

	user, err := entity.NewUser(input.Name, input.LastName, input.Email)
	if err != nil {
		return err
	}

	userModel := &entity.User{
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
	}

	if err := u.userRepository.CreateUser(ctx, userModel); err != nil {
		return err
	}

	return nil
}
