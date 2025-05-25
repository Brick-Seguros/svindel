package entitystrat

import (
	"context"
	"fmt"
	"strings"
	"time"

	"svindel/internal/shared"
	"svindel/pkg/idgen"

	"github.com/sashabaranov/go-openai"
)

type EntityEvalStrategy struct {
	llm openai.Client
}

func New(openAiKey string) *EntityEvalStrategy {

	llm := openai.NewClient(openAiKey)

	return &EntityEvalStrategy{
		llm: *llm,
	}
}

func (s *EntityEvalStrategy) GetCriteria() shared.EvaluationStrategy {
	return shared.EntityStrategy
}

func (s *EntityEvalStrategy) Evaluate(req shared.EvaluationRequest) (shared.EvaluationResult, error) {
	prompt := s.buildPrompt(req)

	raw, err := s.llm.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT4oMini,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	})
	if err != nil {
		return shared.EvaluationResult{}, err
	}

	parsed, err := parseResponse(raw.Choices[0].Message.Content)
	if err != nil {
		return shared.EvaluationResult{}, err
	}

	results := []shared.EvaluationResultItem{}

	results = append(results, shared.EvaluationResultItem{
		CriteriaType: shared.FieldCriteria,
		CriteriaName: "extracted_document_type",
		Value:        parsed.ExpectedDocumentType,
	})

	results = append(results, shared.EvaluationResultItem{
		CriteriaType: shared.FieldCriteria,
		CriteriaName: "extracted_document_value",
		Value:        parsed.ExtractedDocumentValue,
	})

	results = append(results, shared.EvaluationResultItem{
		CriteriaType: shared.FieldCriteria,
		CriteriaName: "expected_document_type",
		Value:        parsed.ExpectedDocumentType,
	})

	results = append(results, shared.EvaluationResultItem{
		CriteriaType: shared.FieldCriteria,
		CriteriaName: "expected_document_value",
		Value:        parsed.ExpectedDocumentValue,
	})

	results = append(results, shared.EvaluationResultItem{
		CriteriaType: shared.TagCriteria,
		CriteriaName: "tags",
		Value:        strings.Join(parsed.Tags, ", "),
	})

	var rating shared.EvaluationResultRating

	rating = shared.EvaluationResultRatingLow

	if parsed.Passed {
		rating = shared.EvaluationResultRatingHigh
	}

	return shared.EvaluationResult{
		Input:       req,
		Comments:    parsed.Comments,
		Strategy:    shared.EntityStrategy,
		EvaluatedAt: time.Now(),
		Results:     results,
		Rating:      rating,
		ID:          idgen.Generate(),
	}, nil
}

// --- Prompt Builder ---

func (s *EntityEvalStrategy) buildPrompt(req shared.EvaluationRequest) string {
	return fmt.Sprintf(`
You are an evaluator for document entity extraction in a fraud analysis system.

Given:

User Input:
%s

AI Response:
%s

Evaluate whether the AI correctly identified the document type and document value.

Valid document types are: CPF, CNPJ, PLATE, NAME, ADDRESS, EMAIL, PHONE or NONE.

Return JSON only in this format:

{
  "passed": true | false,
  "expected_document_type": "CPF | CNPJ | PLATE | NAME | ADDRESS | EMAIL | PHONE | NONE",
  "expected_document_value": "...",
  "extracted_document_type": "CPF | CNPJ | PLATE | NAME | ADDRESS | EMAIL | PHONE | NONE",
  "extracted_document_value": "...",
  "tags": ["correct", "missing", "wrong_type", "wrong_value"],
  "comments": "..."
}

Do not include anything else besides the JSON.
`, req.UserInput, req.AIResponse)
}
