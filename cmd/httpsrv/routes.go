package httpsrv

import (
	"net/http"

	chat "svindel/internal/chat"

	"github.com/gorilla/mux"
)

func NewRouter(chatService *chat.Service) http.Handler {
	r := mux.NewRouter()

	// REST
	r.HandleFunc("/api/chat", HandleChatMessage(chatService)).Methods("POST")

	// WebSocket
	r.HandleFunc("/ws/chat", HandleWebSocket(chatService))

	// Health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	return r
}
