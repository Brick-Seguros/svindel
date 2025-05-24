package docext

import (
	"fmt"

	shared "svindel/internal/shared"
)

func BuildReportsSection(reports []shared.ReportResult) string {
	if len(reports) == 0 {
		return "No related reports found.\n"
	}

	result := "Related Reports:\n"

	for _, r := range reports {
		result += fmt.Sprintf(
			"\n --  ID: %s | Title: %s | Document: %s | CreatedAt: %s\n",
			r.ID, r.Name, r.Document, r.CreatedAt,
		)
	}

	return result
}

func BuildResourcesSection(resources []shared.Resource) string {
	if len(resources) == 0 {
		return "No resources found.\n"
	}

	result := "Available Resources:\n"

	for _, r := range resources {
		result += fmt.Sprintf(
			"\n -- ID: %s | Title: %s | HelperText: %s\n",
			r.ID, r.Title, r.HelperText,
		)
	}

	return result
}

func BuildPromptWithContext(
	userMessage string,
	extraction shared.ExtractionResult,
	context shared.ContextResult,
) string {
	if extraction.IsQuestion {
		return userMessage
	}

	return fmt.Sprintf(`

	---------

	Context for this document:
	- Document: %s
	- Type: %s
	- %v
	- %v


	---------

	User Request:
	%s
`,
		extraction.Document,
		extraction.DocumentType,
		BuildReportsSection(context.Reports),
		BuildResourcesSection(context.Resources),
		userMessage,
	)
}
