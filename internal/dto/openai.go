package dto

// MessageCompletion representa a estrutura da resposta completa de uma mensagem.
type MessageCompletion struct {
	ID                string
	Object            string
	Created           int64
	Model             string
	Choices           []Choice
	Usage             Usage
	SystemFingerprint string `json:"system_fingerprint"`
}

// Choice representa as escolhas dentro da resposta.
type Choice struct {
	Index        int
	Message      MessageContent
	LogProbs     interface{} `json:"logprobs"` // Usar interface{} pois LogProbs é null no exemplo.
	FinishReason string      `json:"finish_reason"`
}

// MessageContent representa o conteúdo da mensagem dentro de uma escolha.
type MessageContent struct {
	Role    string
	Content string
}

// Usage representa o uso de tokens da mensagem.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
