package evalinfra

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"svindel/internal/eval"
	"svindel/internal/shared"
)

type EvalRepository struct {
}

// GetByAIResponse implements eval.Repository.
func (e *EvalRepository) GetByAIResponse(aiResponseID string) ([]shared.EvaluationResult, error) {
	panic("unimplemented")
}

// Save implements eval.Repository.
func (e *EvalRepository) Save(result shared.EvaluationResult) error {
	// save to tmp/eval/
	fmt.Println("Saving evaluation result to", result.ID)

	dir := "tmp/eval/" + string(result.Strategy)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}

	filePath := fmt.Sprintf("%s/%s.json", dir, result.ID)

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, jsonData, 0644)

	if err != nil {
		return fmt.Errorf("failed to save evaluation result to %s: %w", filePath, err)
	}

	log.Println("Saved evaluation result to", filePath)
	return nil
}

func New() eval.Repository {
	return &EvalRepository{}
}
