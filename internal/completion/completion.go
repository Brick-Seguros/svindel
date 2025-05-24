package completion

import (
	"context"
	"fmt"
	"io"
	"time"

	openai_sdk "github.com/sashabaranov/go-openai"

	shared "svindel/internal/shared"
	idgen "svindel/pkg/idgen"
)

// --- Client Struct ---

type Client struct {
	api *openai_sdk.Client
}

func New(apiKey string) *Client {
	return &Client{
		api: openai_sdk.NewClient(apiKey),
	}
}

// --- Sync Completion ---

func (c *Client) GenerateResponse(chatContext shared.Chat) (shared.AIResponse, error) {
	messages := convertToOpenAIMessages(chatContext)

	fmt.Println("messages", messages)

	resp, err := c.api.CreateChatCompletion(
		context.Background(),
		openai_sdk.ChatCompletionRequest{
			Model:    openai_sdk.GPT4o,
			Messages: messages,
		},
	)
	if err != nil {
		return shared.AIResponse{}, err
	}

	raw := resp.Choices[0].Message.Content

	aiMessages, err := parseAIResponse(raw)
	if err != nil {
		return shared.AIResponse{}, fmt.Errorf("parse error: %w", err)
	}

	return shared.AIResponse{
		ID:         idgen.Generate(),
		Messages:   aiMessages,
		CreatedAt:  time.Now(),
		ReceivedAt: time.Now(),
	}, nil
}

// --- Streaming Completion ---

func (c *Client) StreamResponse(
	chatContext shared.Chat,
	onMessage func(shared.AIMessage),
	onDone func(),
	onError func(error),
) error {
	messages := convertToOpenAIMessages(chatContext)

	stream, err := c.api.CreateChatCompletionStream(
		context.Background(),
		openai_sdk.ChatCompletionRequest{
			Model:    openai_sdk.GPT4o,
			Messages: messages,
			Stream:   true,
		},
	)
	if err != nil {
		onError(err)
		return err
	}
	defer stream.Close()

	var currentContent string

	for {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				// Finished
				aiMessages, parseErr := parseAIResponse(currentContent)
				if parseErr != nil {
					onError(parseErr)
					return parseErr
				}

				for _, m := range aiMessages {
					onMessage(m)
				}

				onDone()
				return nil
			}
			onError(err)
			return err
		}

		delta := resp.Choices[0].Delta.Content
		if delta != "" {
			currentContent += delta
		}
	}
}
