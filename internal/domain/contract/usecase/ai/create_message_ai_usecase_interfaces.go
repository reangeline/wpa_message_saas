package usecase

import (
	"context"
)

type CreateMessageAIUseCaseInterface interface {
	Execute(ctx context.Context, input string) (string, error)
}
