package completion

import (
	"encoding/json"
	"fmt"

	shared "svindel/internal/shared"
)

type openAIResponseFormat struct {
	Messages []struct {
		Type    string   `json:"type"`
		Content []string `json:"content"`
	} `json:"messages"`
}

func parseAIResponse(rawText string) ([]shared.AIMessage, error) {
	var parsed openAIResponseFormat

	err := json.Unmarshal([]byte(rawText), &parsed)
	if err != nil {
		return nil, fmt.Errorf("failed to parse AI response JSON: %w", err)
	}

	var messages []shared.AIMessage

	for _, m := range parsed.Messages {
		messageType := shared.AIMessageType(m.Type)

		// Validate the type
		switch messageType {
		case shared.AIMessageTypeText,
			shared.AIMessageTypeReportShortcut,
			shared.AIMessageTypeAgentTrigger,
			shared.AIMessageTypeResourceSelector:
			// Valid types
		default:
			return nil, fmt.Errorf("invalid message type: %s", m.Type)
		}

		messages = append(messages, shared.AIMessage{
			Type:    messageType,
			Content: m.Content,
		})
	}

	return messages, nil
}
