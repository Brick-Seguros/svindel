package reportinfra

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	shared "svindel/internal/shared"
)

type ReportAPIRepository struct {
	baseURL string
	client  *http.Client
	token   string
}

func NewReportAPIRepository(baseURL, token string, client *http.Client) *ReportAPIRepository {
	return &ReportAPIRepository{
		baseURL: baseURL,
		client:  client,
		token:   token,
	}
}

func (r *ReportAPIRepository) FindReportsByDocument(doc string, docType string) ([]shared.ReportResult, error) {
	// Build URL
	u, _ := url.Parse(fmt.Sprintf("%s/v1/record", r.baseURL))
	q := u.Query()
	q.Set("document", doc)
	q.Set("page", "1")
	q.Set("take", "10")
	u.RawQuery = q.Encode()

	// Build Request
	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.token))
	req.Header.Set("Content-Type", "application/json")

	// Execute Request
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Handle non-200
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status from Brick API: %s", resp.Status)
	}

	// Parse Response
	var res struct {
		Data []struct {
			ID        string `json:"id"`
			Name      string `json:"name"`
			Document  string `json:"document"`
			CreatedAt string `json:"created_at"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	var reports []shared.ReportResult

	// for _, r := range res.Data {
	// 	reports = append(reports, shared.ReportResult{
	// 		ID:        r.ID,
	// 		Name:      r.Name,
	// 		Document:  r.Document,
	// 		CreatedAt: r.CreatedAt,
	// 	})
	// }

	// Catching only the first report for now
	if len(res.Data) > 0 {
		reports = append(reports, res.Data[0])
	}

	return reports, nil
}

func (r *ReportAPIRepository) FindReportByID(id string) (shared.ReportResult, error) {
	// Build URL
	url := fmt.Sprintf("%s/v1/record/%s", r.baseURL, id)

	// Build Request
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.token))
	req.Header.Set("Content-Type", "application/json")

	// Execute Request
	resp, err := r.client.Do(req)
	if err != nil {
		return shared.ReportResult{}, err
	}
	defer resp.Body.Close()

	// Handle non-200
	if resp.StatusCode != http.StatusOK {
		return shared.ReportResult{}, fmt.Errorf("unexpected status from Brick API: %s", resp.Status)
	}

	// Parse Response
	var rData struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Document  string `json:"document"`
		CreatedAt string `json:"created_at"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&rData); err != nil {
		return shared.ReportResult{}, err
	}

	return shared.ReportResult{
		ID:        rData.ID,
		Name:      rData.Name,
		Document:  rData.Document,
		CreatedAt: rData.CreatedAt,
	}, nil
}
