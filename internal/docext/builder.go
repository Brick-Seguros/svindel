package docext

import (
	"fmt"

	shared "svindel/internal/shared"
)

func BuildPromptWithContext(
	userMessage string,
	extraction shared.ExtractionResult,
	context ContextResult,
) string {
	if extraction.IsQuestion {
		return userMessage
	}

	return fmt.Sprintf(`
	Context for this document:
	- Document: %s
	- Type: %s
	- Related Reports: %v
	- Available Resources: %v

	User Request:
	%s
`,
		extraction.Document,
		extraction.DocumentType,
		context.Reports,
		context.Resources,
		userMessage,
	)
}
