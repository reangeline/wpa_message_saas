package http

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-chi/chi/v5"
	"github.com/reangeline/micro_saas/internal/dependency_injection"
	"github.com/reangeline/micro_saas/internal/infra/http/route"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
)

var chiLambda *chiadapter.ChiLambda

type ServerLambda struct {
	svc *dynamodb.DynamoDB
}

func NewServerLambda(svc *dynamodb.DynamoDB) *ServerLambda {
	return &ServerLambda{
		svc,
	}
}

func (sl *ServerLambda) ServerHttp() {
	lambda.Start(sl.HandlerRequest)
}

func (sl *ServerLambda) HandlerRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	router := chi.NewRouter()

	route.InitializeMiddlewares(router)

	iu, err := dependency_injection.InitializeUser(sl.svc)
	if err != nil {
		log.Fatalf("failed to initialize user controller: %v", err)
	}

	cwa, err := dependency_injection.InitializeCreateMessageWhatsApp()
	if err != nil {
		log.Fatalf("failed to initialize user controller: %v", err)
	}

	route.InitializeUserRoutes(iu, router)
	route.InitializeWhatsAppRoutes(cwa, router)

	chiLambda = chiadapter.New(router)

	return chiLambda.Proxy(request)
}
