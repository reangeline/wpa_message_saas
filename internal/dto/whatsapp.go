package dto

// Message from a user
type Message struct {
	From string `json:"from"`
	Text Text   `json:"text"`
}

// Text content of the message
type Text struct {
	Body string `json:"body"`
}

// Metadata associated with the message
type Metadata struct {
	PhoneNumberID string `json:"phone_number_id"`
}

// Value structure inside changes
type Value struct {
	Messages []Message `json:"messages"`
	Metadata Metadata  `json:"metadata"`
}

// Change occurring in an entry
type Change struct {
	Value Value `json:"value"`
}

// Entry in the payload
type Entry struct {
	Changes []Change `json:"changes"`
}

// MessagePayload struct holding the whole message structure
type MessagePayload struct {
	Object string  `json:"object"`
	Entry  []Entry `json:"entry"`
}

// Define a struct for the language details
type Language struct {
	Code string `json:"code"`
}

// Define a struct for the template details
type Template struct {
	Name     string   `json:"name"`
	Language Language `json:"language"`
}

// Define a struct for the message data
type MessageDataReply struct {
	MessagingProduct string   `json:"messaging_product"`
	To               string   `json:"to"`
	Type             string   `json:"type"`
	Template         Template `json:"template"`
}
