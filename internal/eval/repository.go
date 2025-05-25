package eval

import "svindel/internal/shared"

type Repository interface {
	Save(result shared.EvaluationResult) error
	GetByAIResponse(aiResponseID string) ([]shared.EvaluationResult, error)
}
