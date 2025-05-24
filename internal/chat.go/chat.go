package chat

import (
	"time"

	docext "svindel/internal/docext"
	shared "svindel/internal/shared"
	idgen "svindel/pkg/idgen"
)

type Service struct {
	ai     AIProvider
	docExt *docext.DocExt
}

func New(ai AIProvider, docExt *docext.DocExt) *Service {
	return &Service{
		ai:     ai,
		docExt: docExt,
	}
}

// --- Chat Lifecycle ---
func (s *Service) NewChat() *shared.Chat {
	return &shared.Chat{
		ID:           idgen.Generate(),
		UserMessages: []shared.UserMessage{},
		AIResponses:  []shared.AIResponse{},
		CreatedAt:    time.Now(),
	}
}

func (s *Service) AddUserMessage(chat *shared.Chat, content string) (shared.UserMessage, error) {
	if content == "" {
		return shared.UserMessage{}, ErrEmptyMessageContent
	}

	msg := shared.UserMessage{
		ID:         idgen.Generate(),
		Content:    content,
		CreatedAt:  time.Now(),
		ReceivedAt: time.Now(),
	}

	chat.UserMessages = append(chat.UserMessages, msg)
	return msg, nil
}

// --- Sync AI Response ---
func (s *Service) GenerateResponse(chat *shared.Chat, userInput string) (shared.AIResponse, error) {
	// Add user message to chat history
	_, err := s.AddUserMessage(chat, userInput)
	if err != nil {
		return shared.AIResponse{}, err
	}

	augPrompt, _, _ := s.docExt.Process(userInput)

	augmentedChat := *chat
	augmentedChat.UserMessages = append(augmentedChat.UserMessages, shared.UserMessage{
		ID:         idgen.Generate(),
		Content:    augPrompt,
		CreatedAt:  time.Now(),
		ReceivedAt: time.Now(),
	})

	resp, err := s.ai.GenerateResponse(augmentedChat)
	if err != nil {
		return shared.AIResponse{}, err
	}

	chat.AIResponses = append(chat.AIResponses, resp)

	return resp, nil
}

// --- Stream AI Response  ---
func (s *Service) StreamAIResponse(
	chat *shared.Chat,
	onMessage func(shared.AIMessage),
	onDone func(),
	onError func(error),
) error {
	return s.ai.StreamResponse(*chat, onMessage, onDone, onError)
}

// --- Feedback ---
func (s *Service) RateAIResponse(chat *shared.Chat, responseID string, liked bool) error {
	for i := range chat.AIResponses {
		if chat.AIResponses[i].ID == responseID {
			chat.AIResponses[i].Liked = &liked
			return nil
		}
	}
	return ErrResponseNotFound
}
