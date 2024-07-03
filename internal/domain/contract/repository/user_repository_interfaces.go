package contract

import (
	"context"

	"github.com/reangeline/micro_saas/internal/domain/entity"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *entity.User) error
	FindAll() ([]*entity.User, error)
	FindByUserEmail(email string) (*entity.User, error)
	UpdateByEmail(input *entity.User) (*entity.User, error)
}
