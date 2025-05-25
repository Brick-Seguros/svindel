package eval

import (
	"log"
	"svindel/internal/shared"
	"sync"
)

type Evaluator struct {
	strategies []shared.Strategy
	repository Repository
}

func New(strategies []shared.Strategy, repository Repository) *Evaluator {
	return &Evaluator{
		strategies: strategies,
		repository: repository,
	}
}

func (e *Evaluator) EvaluateAsync(req shared.EvaluationRequest) {
	go e.evaluate(req)
}

func (e *Evaluator) evaluate(req shared.EvaluationRequest) {
	var wg sync.WaitGroup

	for _, strategy := range e.strategies {
		wg.Add(1)

		go func(s shared.Strategy) {
			defer wg.Done()

			res, err := s.Evaluate(req)
			if err != nil {
				log.Println("Evaluation error:", err)
				return
			}

			if err := e.repository.Save(res); err != nil {
				log.Println("Failed to save evaluation:", err)
			}
		}(strategy)
	}

	wg.Wait()
}
