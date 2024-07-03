package controller

import (
	"encoding/json"
	"net/http"

	usecase "github.com/reangeline/micro_saas/internal/domain/contract/usecase/wp"
	"github.com/reangeline/micro_saas/internal/dto"
)

type WhatsAppController struct {
	createMessageWhatAppUseCase usecase.CreateMessageWhatsAppUseCaseInterface
}

func NewWhatsAppController(
	createMessageWhatAppUseCase usecase.CreateMessageWhatsAppUseCaseInterface,

) *WhatsAppController {
	return &WhatsAppController{
		createMessageWhatAppUseCase: createMessageWhatAppUseCase,
	}
}

func (wa *WhatsAppController) CreateMessageWhatsApp(w http.ResponseWriter, r *http.Request) {

	var input *dto.MessagePayload
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	ctx := r.Context()

	if input.Object != "" {
		wa.createMessageWhatAppUseCase.Execute(ctx, input)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}

func (wa *WhatsAppController) VerifyTokenWhatsApp(w http.ResponseWriter, r *http.Request) {
	verifyToken := "TESTE"
	mode := r.URL.Query().Get("hub.mode")
	token := r.URL.Query().Get("hub.verify_token")
	challenge := r.URL.Query().Get("hub.challenge")

	if mode == "subscribe" && token == verifyToken {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(challenge))
		return
	}

	http.Error(w, "Forbidden", http.StatusForbidden)
}
