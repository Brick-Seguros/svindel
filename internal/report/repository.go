package report

import shared "svindel/internal/shared"

type Repository interface {
	FindReportsByDocument(doc string, docType string) ([]shared.ReportResult, error)
	FindReportByID(id string) (shared.ReportResult, error)
}
