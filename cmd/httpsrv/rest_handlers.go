package httpsrv

import (
	"encoding/json"
	"net/http"

	chat "svindel/internal/chat.go"
)

type ChatRequest struct {
	Message string `json:"message"`
}

type ChatResponse struct {
	Reply []string `json:"reply"`
}

func HandleChatMessage(service *chat.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ChatRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		chatSession := service.NewChat()

		_, err := service.AddUserMessage(chatSession, req.Message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp, err := service.GenerateResponse(chatSession, req.Message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var texts []string
		for _, msg := range resp.Messages {
			texts = append(texts, msg.Content...)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ChatResponse{
			Reply: texts,
		})
	}
}
