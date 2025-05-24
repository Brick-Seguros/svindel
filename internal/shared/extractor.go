package shared

type DocExt interface {
	Process(userMessage string) (augmentedPrompt string, extraction ExtractionResult, context ContextResult)
}
