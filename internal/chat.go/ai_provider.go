package chat

import shared "svindel/internal/shared"

type AIProvider interface {
	GenerateResponse(chat shared.Chat) (shared.AIResponse, error)

	StreamResponse(
		chat shared.Chat,
		onMessage func(shared.AIMessage),
		onDone func(),
		onError func(error),
	) error
}
