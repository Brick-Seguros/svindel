package chat

import "errors"

var (
	ErrEmptyMessageContent = errors.New("user message content cannot be empty")
	ErrResponseNotFound    = errors.New("AI response not found")
	ErrChatNotFound        = errors.New("chat not found")
)
