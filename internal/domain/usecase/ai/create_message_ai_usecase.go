package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type CreateMessageAIUseCase struct {
}

func NewCreateMessageAIUseCase() *CreateMessageAIUseCase {
	return &CreateMessageAIUseCase{}
}

func (mai *CreateMessageAIUseCase) Execute(ctx context.Context, input string) (string, error) {

	apiKey := ""

	url := "https://api.openai.com/v1/chat/completions"

	loja := map[string]string{
		"nome":     "Loja Exemplo",
		"endereco": "Rua Exemplo, 123",
		"contato":  "exemplo@loja.com",
	}

	produtos := []map[string]string{
		{"nome": "Produto A", "preco": "R$ 10,00"},
		{"nome": "Produto B", "preco": "R$ 20,00"},
		{"nome": "Produto C", "preco": "R$ 30,00"},
	}

	// usuario := map[string]string{
	// 	"nome":  "Renato",
	// 	"email": "renato@exemplo.com",
	// }

	storeInfo := fmt.Sprintf("Store: %s, Address: %s, Contact: %s", loja["nome"], loja["endereco"], loja["contato"])
	productsInfo := "Products:\n"
	for _, product := range produtos {
		productsInfo += fmt.Sprintf("- %s: %s\n", product["nome"], product["preco"])
	}

	content := fmt.Sprintf("%s\n%s\n\n%s", storeInfo, productsInfo, input)

	payload := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]interface{}{
			{
				"role":    "system",
				"content": "You are a helpful assistant.",
			},
			{
				"role":    "user",
				"content": content,
			},
		},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error occurred during marshaling. Error: %s", err.Error())
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Error occurred during request creation. Error: %s", err.Error())
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error occurred during request execution. Error: %s", err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error occurred during reading the response body. Error: %s", err.Error())
	}

	return string(body), nil

}
