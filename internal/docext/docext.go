package docext

import (
	shared "svindel/internal/shared"
)

type DocExt struct {
	Extractor Extractor
	Retriever Retriever
}

func New(
	extractor Extractor,
	retriever Retriever,
) *DocExt {
	return &DocExt{
		Extractor: extractor,
		Retriever: retriever,
	}
}

func (d *DocExt) Process(
	userMessage string,
) (augmentedPrompt string, extraction shared.ExtractionResult, context ContextResult) {
	// Step 1 — Extract document from prompt
	extraction = d.Extractor.Extract(userMessage)

	// Step 2 — Retrieve context if it's not a question
	if !extraction.IsQuestion {
		context = d.Retriever.Retrieve(extraction.Document, extraction.DocumentType)
		augmentedPrompt = BuildPromptWithContext(userMessage, extraction, context)
	} else {
		augmentedPrompt = userMessage
	}

	return
}
