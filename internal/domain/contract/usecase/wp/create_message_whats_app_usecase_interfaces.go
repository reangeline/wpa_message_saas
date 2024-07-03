package usecase

import (
	"context"

	"github.com/reangeline/micro_saas/internal/dto"
)

type CreateMessageWhatsAppUseCaseInterface interface {
	Execute(ctx context.Context, input *dto.MessagePayload) error
}
