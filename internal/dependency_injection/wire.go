//go:build wireinject
// +build wireinject

package dependency_injection

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/wire"
	repo_interface "github.com/reangeline/micro_saas/internal/domain/contract/repository"
	ai_usecase_inteface "github.com/reangeline/micro_saas/internal/domain/contract/usecase/ai"
	user_usecase_inteface "github.com/reangeline/micro_saas/internal/domain/contract/usecase/user"
	wp_usecase_inteface "github.com/reangeline/micro_saas/internal/domain/contract/usecase/wp"

	ai_usecase "github.com/reangeline/micro_saas/internal/domain/usecase/ai"
	user_usecase "github.com/reangeline/micro_saas/internal/domain/usecase/user"
	wp_usecase "github.com/reangeline/micro_saas/internal/domain/usecase/wp"

	repository "github.com/reangeline/micro_saas/internal/infra/database/repository"

	"github.com/reangeline/micro_saas/internal/presentation/controller"
)

var setCreateUserUseCaseDependency = wire.NewSet(
	user_usecase.NewCreateUserUseCase,
	wire.Bind(new(user_usecase_inteface.CreateUserUseCaseInterface), new(*user_usecase.CreateUserUseCase)),
)

var setUserRepositoryDependency = wire.NewSet(
	repository.NewUserRepository,
	wire.Bind(new(repo_interface.UserRepositoryInterface), new(*repository.UserRepository)),
)

var setCreateMessageWhatsAppUseCaseDependency = wire.NewSet(
	wp_usecase.NewCreateMessageWhatsAppUseCase,
	wire.Bind(new(wp_usecase_inteface.CreateMessageWhatsAppUseCaseInterface), new(*wp_usecase.CreateMessageWhatsAppUseCase)),
)

var setCreateMessageAIUseCaseDependency = wire.NewSet(
	ai_usecase.NewCreateMessageAIUseCase,
	wire.Bind(new(ai_usecase_inteface.CreateMessageAIUseCaseInterface), new(*ai_usecase.CreateMessageAIUseCase)),
)

func InitializeUser(vc *dynamodb.DynamoDB) (*controller.UserController, error) {
	wire.Build(
		setUserRepositoryDependency,
		setCreateUserUseCaseDependency,
		controller.NewUserController,
	)
	return &controller.UserController{}, nil
}

func InitializeCreateMessageWhatsApp() (*controller.WhatsAppController, error) {
	wire.Build(
		setCreateMessageAIUseCaseDependency,
		setCreateMessageWhatsAppUseCaseDependency,
		controller.NewWhatsAppController,
	)
	return &controller.WhatsAppController{}, nil
}
