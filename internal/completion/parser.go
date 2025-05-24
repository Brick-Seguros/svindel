package completion

import (
	"encoding/json"
	"fmt"

	shared "svindel/internal/shared"
)

type openAIResponseFormat struct {
	Messages []struct {
		Type     string `json:"type"`
		Text     string `json:"text"`
		Shortcut struct {
			ID        string `json:"id"`
			Title     string `json:"title"`
			Document  string `json:"document"`
			CreatedAt string `json:"createdAt"`
		} `json:"shortcut"`
		Resources []struct {
			ID         string `json:"id"`
			Title      string `json:"title"`
			HelperText string `json:"helperText"`
		} `json:"resources"`
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
			Type: messageType,
			Text: m.Text,
			Shortcut: shared.Shortcut{
				ID:        m.Shortcut.ID,
				Title:     m.Shortcut.Title,
				Document:  m.Shortcut.Document,
				CreatedAt: m.Shortcut.CreatedAt,
			},
			Resources: func() []shared.Resource {
				resources := []shared.Resource{}
				for _, r := range m.Resources {
					resources = append(resources, shared.Resource{ID: r.ID, Title: r.Title, HelperText: r.HelperText})
				}
				return resources
			}(),
		})
	}

	return messages, nil
}
