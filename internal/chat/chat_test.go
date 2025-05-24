package chat

import (
	"testing"
	"time"

	"svindel/internal/shared"
)

// ---- Mocks ----

// Mock AIProvider
type MockAIProvider struct{}

func (m *MockAIProvider) GenerateResponse(chat shared.Chat) (shared.AIResponse, error) {
	return shared.AIResponse{
		ID: "ai-response-123",
		Messages: []shared.AIMessage{
			{
				Type: shared.AIMessageTypeText,
				Text: "This is a test response.",
			},
		},
		CreatedAt:  time.Now(),
		ReceivedAt: time.Now(),
	}, nil
}

func (m *MockAIProvider) StreamResponse(
	chat shared.Chat,
	onMessage func(shared.AIMessage),
	onDone func(),
	onError func(error),
) error {
	onMessage(shared.AIMessage{
		Type: shared.AIMessageTypeText,
		Text: "Streamed message part",
	})
	onDone()
	return nil
}

// Mock DocExt
type MockDocExt struct{}

func (m *MockDocExt) Process(input string) (string, shared.ExtractionResult, shared.ContextResult) {
	return input + " + context", shared.ExtractionResult{}, shared.ContextResult{}
}

// ---- Tests ----

func TestNewChat(t *testing.T) {
	svc := New(&MockAIProvider{}, &MockDocExt{})
	chat := svc.NewChat()

	if chat.ID == "" {
		t.Error("Expected chat to have an ID")
	}
	if chat.CreatedAt.IsZero() {
		t.Error("Expected CreatedAt to be set")
	}
	if len(chat.UserMessages) != 0 {
		t.Error("Expected no user messages initially")
	}
	if len(chat.AIResponses) != 0 {
		t.Error("Expected no AI responses initially")
	}
}

func TestAddUserMessage(t *testing.T) {
	svc := New(&MockAIProvider{}, &MockDocExt{})
	chat := svc.NewChat()

	msg, err := svc.AddUserMessage(chat, "Hello")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if msg.Content != "Hello" {
		t.Errorf("Expected content 'Hello', got '%s'", msg.Content)
	}

	if len(chat.UserMessages) != 1 {
		t.Errorf("Expected 1 user message, got %d", len(chat.UserMessages))
	}
}

func TestGenerateResponse(t *testing.T) {
	svc := New(&MockAIProvider{}, &MockDocExt{})
	chat := svc.NewChat()

	resp, err := svc.GenerateResponse(chat, "Hello")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if resp.ID != "ai-response-123" {
		t.Errorf("Expected response ID 'ai-response-123', got '%s'", resp.ID)
	}

	if len(resp.Messages) != 1 {
		t.Errorf("Expected 1 AI message, got %d", len(resp.Messages))
	}

	if resp.Messages[0].Text != "This is a test response." {
		t.Errorf("Unexpected AI message text: %s", resp.Messages[0].Text)
	}

	if len(chat.UserMessages) != 1 {
		t.Errorf("Expected 1 user message, got %d", len(chat.UserMessages))
	}

	if len(chat.AIResponses) != 1 {
		t.Errorf("Expected 1 AI response, got %d", len(chat.AIResponses))
	}
}

func TestRateAIResponse(t *testing.T) {
	svc := New(&MockAIProvider{}, &MockDocExt{})
	chat := svc.NewChat()

	// Add fake response
	resp := shared.AIResponse{
		ID: "ai-response-123",
	}
	chat.AIResponses = append(chat.AIResponses, resp)

	// Rate it
	err := svc.RateAIResponse(chat, "ai-response-123", true)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if chat.AIResponses[0].Liked == nil || *chat.AIResponses[0].Liked != true {
		t.Errorf("Expected Liked to be true, got %v", chat.AIResponses[0].Liked)
	}
}

func TestRateAIResponse_NotFound(t *testing.T) {
	svc := New(&MockAIProvider{}, &MockDocExt{})
	chat := svc.NewChat()

	err := svc.RateAIResponse(chat, "non-existent", true)
	if err == nil {
		t.Error("Expected error when rating a non-existent response")
	}
}
