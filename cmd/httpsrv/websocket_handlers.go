package httpsrv

import (
	"log"
	"net/http"
	"time"

	chat "svindel/internal/chat"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 🔥 You can add security rules here (e.g., CORS, auth).
	},
}

func HandleWebSocket(service *chat.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("❌ WebSocket upgrade error:", err)
			return
		}
		defer conn.Close()

		// 🔥 Assign a unique connection ID
		connectionID := uuid.NewString()

		log.Printf("🟢 WebSocket connected | ID: %s | Remote: %s | Time: %s",
			connectionID, r.RemoteAddr, time.Now().Format(time.RFC3339))

		defer log.Printf("🔴 WebSocket disconnected | ID: %s | Remote: %s | Time: %s",
			connectionID, r.RemoteAddr, time.Now().Format(time.RFC3339))

		chatSession := service.NewChat()

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("⚠️  Read error | ID: %s | Error: %v", connectionID, err)
				break
			}

			log.Printf("📩 Message received | ID: %s | Message: %s", connectionID, string(message))

			// Run Doc Extraction + Prompt Augmentation
			response, err := service.GenerateResponse(chatSession, string(message))
			if err != nil {
				conn.WriteJSON(map[string]string{"error": err.Error()})
				log.Printf("❌ Error generating response | ID: %s | Error: %v", connectionID, err)
				continue
			}

			// Streaming back AI Messages (multi-message response)
			for _, msg := range response.Messages {
				conn.WriteJSON(msg)
				log.Printf("📤 Message sent | ID: %s | Type: %s", connectionID, msg.Type)
			}

			// "done" event
			conn.WriteJSON(map[string]string{"event": "done"})
			log.Printf("✅ Done event sent | ID: %s", connectionID)
		}
	}
}
