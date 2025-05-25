package relevancestrat

import (
	"context"
	"fmt"
	"strings"
	"time"

	"svindel/internal/shared"
	"svindel/pkg/idgen"

	"github.com/sashabaranov/go-openai"
)

type RelevanceEvalStrategy struct {
	llm openai.Client
}

func New(openAiKey string) *RelevanceEvalStrategy {
	llm := openai.NewClient(openAiKey)

	return &RelevanceEvalStrategy{
		llm: *llm,
	}
}

func (s *RelevanceEvalStrategy) GetCriteria() shared.EvaluationStrategy {
	return shared.RelevanceStrategy
}

func (s *RelevanceEvalStrategy) Evaluate(req shared.EvaluationRequest) (shared.EvaluationResult, error) {
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
		CriteriaType: shared.ScoreCriteria,
		CriteriaName: "relevance_score",
		Value:        fmt.Sprintf("%d", parsed.RelevanceScore),
	})

	results = append(results, shared.EvaluationResultItem{
		CriteriaType: shared.TagCriteria,
		CriteriaName: "tags",
		Value:        strings.Join(parsed.Tags, ", "),
	})

	var rating shared.EvaluationResultRating

	rating = shared.EvaluationResultRatingLow

	if parsed.RelevanceScore >= 8 {
		rating = shared.EvaluationResultRatingHigh
	} else if parsed.RelevanceScore >= 5 {
		rating = shared.EvaluationResultRatingMedium
	}

	return shared.EvaluationResult{
		Input:       req,
		Comments:    parsed.Comments,
		Strategy:    shared.RelevanceStrategy,
		EvaluatedAt: time.Now(),
		Results:     results,
		Rating:      rating,
		ID:          idgen.Generate(),
	}, nil
}

// --- Prompt Builder ---

func (s *RelevanceEvalStrategy) buildPrompt(req shared.EvaluationRequest) string {
	return fmt.Sprintf(`
You are an evaluator for AI response relevance in a fraud analysis system.

Given:

- User Input:
%s

- Context (including reports, resources, and available data):
%s

- AI Response:
%s

Evaluate whether the AI response is relevant to the user input and the context of fraud detection.

Return JSON ONLY in this format:

{
  "relevance_score": 0-10,
  "passed": true | false,
  "tags": ["highly_relevant", "partially_relevant", "irrelevant"],
  "comments": "..."
}

Do NOT include anything else beyond the JSON.
`,
		req.UserInput,
		req.Context,
		req.AIResponse,
	)
}
