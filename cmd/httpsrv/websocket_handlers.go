package httpsrv

import (
	"log"
	"net/http"

	chat "svindel/internal/chat.go"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins (add security rules here)
	},
}

func HandleWebSocket(service *chat.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("WebSocket upgrade error:", err)
			return
		}
		defer conn.Close()

		chatSession := service.NewChat()

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("WebSocket read error:", err)
				break
			}

			// 🔥 Run Doc Extraction + Prompt Augmentation
			response, err := service.GenerateResponse(chatSession, string(message))
			if err != nil {
				conn.WriteJSON(map[string]string{"error": err.Error()})
				continue
			}

			// 🔥 Stream back AI Messages (multi-message response)
			for _, msg := range response.Messages {
				conn.WriteJSON(msg)
			}

			// 🔥 Optionally send a "done" event
			conn.WriteJSON(map[string]string{"event": "done"})
		}
	}
}
