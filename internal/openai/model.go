package openai

type Response struct {
	Data    []Message `json:"data" validate:"required,dive"`
	FirstID string    `json:"first_id" validate:"required"`
	HasMore bool      `json:"has_more"`
	LastID  string    `json:"last_id" validate:"required"`
	Object  string    `json:"object" validate:"required,eq=list"`
}

type Message struct {
	AssistantID *string   `json:"assistant_id" validate:"omitempty"`
	Attachments []string  `json:"attachments"`
	Content     []Content `json:"content" validate:"required,dive"`
	CreatedAt   int64     `json:"created_at" validate:"required"`
	ID          string    `json:"id" validate:"required"`
	Metadata    Metadata  `json:"metadata"`
	Object      string    `json:"object" validate:"required,eq=thread.message"`
	Role        string    `json:"role" validate:"required,oneof=assistant user"`
	RunID       *string   `json:"run_id" validate:"omitempty"`
	ThreadID    string    `json:"thread_id" validate:"required"`
}

type Content struct {
	Text TextContent `json:"text" validate:"required"`
	Type string      `json:"type" validate:"required,eq=text"`
}

type TextContent struct {
	Annotations []string `json:"annotations"`
	Value       string   `json:"value" validate:"required"`
}

type Metadata map[string]interface{}
