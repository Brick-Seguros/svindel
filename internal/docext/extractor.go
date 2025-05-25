package docext

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	shared "svindel/internal/shared"

	openai_sdk "github.com/sashabaranov/go-openai"
)

type Extractor interface {
	Extract(input string) shared.ExtractionResult
}

type OpenAiExtractor struct {
	openai      *openai_sdk.Client
	evalService shared.Evaluator
}

func NewOpenAiExtractor(openaiKey string,
	evalService shared.Evaluator,
) Extractor {
	client := openai_sdk.NewClient(openaiKey)
	return &OpenAiExtractor{openai: client, evalService: evalService}
}

func (e *OpenAiExtractor) Extract(input string) shared.ExtractionResult {
	// input := strings.ToLower(input)

	// --- Regex CPF ---
	cpfPattern := regexp.MustCompile(`\b\d{3}\.?\d{3}\.?\d{3}-?\d{2}\b`)
	if cpf := cpfPattern.FindString(input); cpf != "" {
		return shared.ExtractionResult{
			Document:     cpf,
			DocumentType: shared.DocTypeCPF,
			IsQuestion:   false,
		}
	}

	// --- Regex CNPJ ---
	cnpjPattern := regexp.MustCompile(`\b\d{2}\.?\d{3}\.?\d{3}/?\d{4}-?\d{2}\b`)
	if cnpj := cnpjPattern.FindString(input); cnpj != "" {
		return shared.ExtractionResult{
			Document:     cnpj,
			DocumentType: shared.DocTypeCNPJ,
			IsQuestion:   false,
		}
	}

	// --- Regex Plate ---
	platePattern := regexp.MustCompile(`[a-zA-Z]{3}[0-9][A-Za-z0-9][0-9]{2}`)
	if plate := platePattern.FindString(input); plate != "" {
		return shared.ExtractionResult{
			Document:     plate,
			DocumentType: shared.DocTypePlate,
			IsQuestion:   false,
		}
	}

	// --- Regex Email ---
	emailPattern := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-z]{2,}`)
	if email := emailPattern.FindString(input); email != "" {
		return shared.ExtractionResult{
			Document:     email,
			DocumentType: shared.DocTypeEmail,
			IsQuestion:   false,
		}
	}

	// --- Fallback to OpenAI ---
	if e.openai == nil {
		return shared.ExtractionResult{
			Document:     "",
			DocumentType: shared.DocTypeNone,
			IsQuestion:   true,
		}
	}

	res := e.extractWithLLM(input)

	e.evalService.EvaluateAsync(shared.EvaluationRequest{
		UserInput:  input,
		AIResponse: res,
	})

	return res
}

func (e *OpenAiExtractor) extractWithLLM(input string) shared.ExtractionResult {
	prompt := buildExtractionPrompt(input)

	resp, err := e.openai.CreateChatCompletion(
		context.Background(),
		openai_sdk.ChatCompletionRequest{
			Model: openai_sdk.GPT4o,
			Messages: []openai_sdk.ChatCompletionMessage{
				{
					Role:    openai_sdk.ChatMessageRoleSystem,
					Content: systemExtractionPrompt,
				},
				{
					Role:    openai_sdk.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		fmt.Println("OpenAI extraction error:", err)
		return shared.ExtractionResult{
			Document:     "",
			DocumentType: shared.DocTypeNone,
			IsQuestion:   true,
		}
	}

	raw := resp.Choices[0].Message.Content

	var parsed struct {
		Document string `json:"document"`
		Type     string `json:"type"`
	}

	err = json.Unmarshal([]byte(raw), &parsed)
	if err != nil {
		fmt.Println("Failed to parse OpenAI JSON:", err)
		return shared.ExtractionResult{
			Document:     "",
			DocumentType: shared.DocTypeNone,
			IsQuestion:   true,
		}
	}

	docType := shared.DocType(strings.ToUpper(parsed.Type))

	// Validate docType
	switch docType {
	case shared.DocTypeCPF, shared.DocTypeCNPJ, shared.DocTypePlate, shared.DocTypeName:
		return shared.ExtractionResult{
			Document:     parsed.Document,
			DocumentType: docType,
			IsQuestion:   false,
		}
	default:
		return shared.ExtractionResult{
			Document:     "",
			DocumentType: shared.DocTypeNone,
			IsQuestion:   true,
		}
	}
}

func buildExtractionPrompt(userInput string) string {
	return fmt.Sprintf(`
		User input: "%s"

		Extract the document and its type (CPF, CNPJ, Name, Plate, Email, Phone, Address). 
		Remember that it can be a CPF, CNPJ, Name, Plate, Email, Phone or a Address.
		If no document is found, classify it as "None".
		Return in JSON with fields: document (string) and type (CPF, CNPJ, Name, Plate, Email, Phone, Address, None).
		Only return valid JSON.`,
		userInput,
	)
}

const systemExtractionPrompt = `
	You are an assistant that extracts structured data from user input.

	Output JSON in the following format:

	{
		"document": string,
		"type": "CPF" | "CNPJ" | "NAME" | "PLATE" | "EMAIL" | "PHONE" | "ADDRESS" | "NONE"
	}

	Rules:
	- "document": the exact document string extracted.
	- "type": the document type or "NONE" if no document is found.
	- If the input is a general question with no document, set "type" to "NONE".
	- Do not include any text besides the JSON.
	- Try your best to recognize a document even if it's not in the exact format.
`
