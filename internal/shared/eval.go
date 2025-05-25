package shared

import "time"

type Evaluator interface {
	EvaluateAsync(req EvaluationRequest)
}

type EvaluationResult struct {
	ID          string
	Input       EvaluationRequest
	Results     []EvaluationResultItem
	Comments    string
	Rating      EvaluationResultRating
	Strategy    EvaluationStrategy
	EvaluatedAt time.Time
}

type EvaluationResultItemCriteria string

const (
	ScoreCriteria   = "score"
	TagCriteria     = "tag"
	FieldCriteria   = "field"
	BooleanCriteria = "boolean"
)

type EvaluationResultItem struct {
	CriteriaType string
	CriteriaName string
	Value        string
}

type EvaluationRequest struct {
	UserInput  string
	Context    string
	AIResponse interface{} // shared.AIResponse or similar
}
