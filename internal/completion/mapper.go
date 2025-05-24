package completion

import (
	openai_sdk "github.com/sashabaranov/go-openai"

	shared "svindel/internal/shared"
)

func convertToOpenAIMessages(c shared.Chat) []openai_sdk.ChatCompletionMessage {
	messages := []openai_sdk.ChatCompletionMessage{}

	messages = append(messages, openai_sdk.ChatCompletionMessage{
		Role:    openai_sdk.ChatMessageRoleSystem,
		Content: systemPrompt,
	})

	for i := 0; i < len(c.UserMessages) && i < len(c.AIResponses); i++ {
		messages = append(messages, openai_sdk.ChatCompletionMessage{
			Role:    openai_sdk.ChatMessageRoleUser,
			Content: c.UserMessages[i].Content,
		})
		// Only append text-based AI responses
		for _, aiMsg := range c.AIResponses[i].Messages {
			if aiMsg.Type == shared.AIMessageTypeText {
				messages = append(messages, openai_sdk.ChatCompletionMessage{
					Role:    openai_sdk.ChatMessageRoleAssistant,
					Content: aiMsg.Text,
				})
			}
		}
	}

	// If user message without response yet
	if len(c.UserMessages) > len(c.AIResponses) {
		last := c.UserMessages[len(c.UserMessages)-1]
		messages = append(messages, openai_sdk.ChatCompletionMessage{
			Role:    openai_sdk.ChatMessageRoleUser,
			Content: last.Content,
		})
	}

	return messages
}
