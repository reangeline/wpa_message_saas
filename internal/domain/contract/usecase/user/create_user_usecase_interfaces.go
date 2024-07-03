package usecase

import (
	"context"

	"github.com/reangeline/micro_saas/internal/dto"
)

type CreateUserUseCaseInterface interface {
	Execute(ctx context.Context, input *dto.UserInput) error
}
