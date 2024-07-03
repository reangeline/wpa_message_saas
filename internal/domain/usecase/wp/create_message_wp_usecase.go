package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	usecase "github.com/reangeline/micro_saas/internal/domain/contract/usecase/ai"
	"github.com/reangeline/micro_saas/internal/dto"
)

type CreateMessageWhatsAppUseCase struct {
	messageAIUsecase usecase.CreateMessageAIUseCaseInterface
}

func NewCreateMessageWhatsAppUseCase(
	messageAIUsecase usecase.CreateMessageAIUseCaseInterface,
) *CreateMessageWhatsAppUseCase {
	return &CreateMessageWhatsAppUseCase{
		messageAIUsecase,
	}
}

func (cwa *CreateMessageWhatsAppUseCase) Execute(ctx context.Context, input *dto.MessagePayload) error {

	if len(input.Entry) > 0 && len(input.Entry[0].Changes) > 0 && len(input.Entry[0].Changes[0].Value.Messages) > 0 {
		// phoneID := input.Entry[0].Changes[0].Value.Metadata.PhoneNumberID
		// from := input.Entry[0].Changes[0].Value.Messages[0].From
		body := input.Entry[0].Changes[0].Value.Messages[0].Text.Body

		url := "https://graph.facebook.com/v19.0/258467377345703/messages"

		reply, err := cwa.messageAIUsecase.Execute(ctx, body)
		if err != nil {
			return err
		}

		var msg dto.MessageCompletion

		// Deserializando a string JSON para a struct MessageCompletion
		err = json.Unmarshal([]byte(reply), &msg)
		if err != nil {
			fmt.Printf("Erro ao deserializar JSON: %s\n", err)
			return err
		}

		// Construct the JSON payload
		messageData := map[string]interface{}{
			"messaging_product": "whatsapp",
			"to":                "5511967700232",
			"text": map[string]interface{}{
				"body": msg.Choices[0].Message.Content,
			},
		}

		// Convert messageData to JSON
		jsonData, err := json.Marshal(messageData)
		if err != nil {
			return err
		}

		// Create the HTTP request
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			return err
		}

		// Set the required headers
		req.Header.Set("Authorization", "Bearer ")
		req.Header.Set("Content-Type", "application/json")

		// Send the request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		log.Println(body)

	}

	return nil
}
