package report

import (
	shared "svindel/internal/shared"
)

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetReportsForDocument(doc string, docType shared.DocType) ([]shared.ReportResult, error) {
	return s.repo.FindReportsByDocument(doc, string(docType))
}
