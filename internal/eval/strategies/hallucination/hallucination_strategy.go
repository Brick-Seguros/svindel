package hallucinationstrat

import (
	"context"
	"fmt"
	"strings"
	"time"

	"svindel/internal/shared"
	"svindel/pkg/idgen"

	"github.com/sashabaranov/go-openai"
)

type HallucinationEvalStrategy struct {
	llm openai.Client
}

func New(openAiKey string) *HallucinationEvalStrategy {
	llm := openai.NewClient(openAiKey)

	return &HallucinationEvalStrategy{
		llm: *llm,
	}
}

func (s *HallucinationEvalStrategy) GetCriteria() shared.EvaluationStrategy {
	return shared.HallucinationStrategy
}

func (s *HallucinationEvalStrategy) Evaluate(req shared.EvaluationRequest) (shared.EvaluationResult, error) {
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
		CriteriaName: "hallucination_risk",
		Value:        parsed.HallucinationRisk,
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
		Strategy:    shared.HallucinationStrategy,
		EvaluatedAt: time.Now(),
		Results:     results,
		Rating:      rating,
		ID:          idgen.Generate(),
	}, nil
}

// --- Prompt Builder ---

func (s *HallucinationEvalStrategy) buildPrompt(req shared.EvaluationRequest) string {
	return fmt.Sprintf(`
You are an evaluator for AI hallucination in a fraud analysis system.

Given:

- User Input:
%s

- Context (including reports, resources, and available data):
%s

- AI Response:
%s

Determine whether the AI response contains any hallucinated content — meaning information, entities, document references, reports, resources, or claims that do NOT exist in the provided context.

Return JSON ONLY in the following format:

{
  "hallucination_risk": "low" | "medium" | "high",
  "passed": true | false,
  "tags": ["correct", "minor_hallucination", "major_hallucination"],
  "comments": "..."
}

Do NOT include anything else besides the JSON.
`,
		req.UserInput,
		req.Context,
		req.AIResponse,
	)
}
