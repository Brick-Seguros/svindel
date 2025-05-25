package shared

type Strategy interface {
	Evaluate(req EvaluationRequest) (EvaluationResult, error)
	GetCriteria() EvaluationStrategy
}

type EvaluationStrategy string

const (
	EntityStrategy        = "entity_recognition"
	HallucinationStrategy = "hallucination"
	RelevanceStrategy     = "relevance"
)

type EvaluationResultRating string

const (
	EvaluationResultRatingLow    = "bad"
	EvaluationResultRatingMedium = "medium"
	EvaluationResultRatingHigh   = "good"
)
