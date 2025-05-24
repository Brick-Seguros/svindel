package report

import (
	shared "svindel/internal/shared"
)

type ReportService struct {
	repo Repository
}

func New(repo Repository) *ReportService {
	return &ReportService{
		repo: repo,
	}
}

func (s *ReportService) GetReportsForDocument(doc string, docType shared.DocType) ([]shared.ReportResult, error) {
	return s.repo.FindReportsByDocument(doc, string(docType))
}
