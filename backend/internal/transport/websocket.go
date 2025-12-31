package transport

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"goTradingg/internal/models"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for simplicity; adjust as needed for security
	},
}

/*
❌ Incorrect mental model
“I need to spawn a goroutine for each WebSocket client”

✅ Correct mental model
Each WebSocket handler is already running in its own goroutine
*/

func (h *HTTPHandler) tradeWS(w http.ResponseWriter, r *http.Request){
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WS upgrade failed:", err)
		return
	}
	defer conn.Close()

	tradeCh := make(chan models.Trade, 10) // buffered channel to avoid blocking
	h.broadcaster.Register(tradeCh)
	defer h.broadcaster.Unregister(tradeCh)

	for trade := range tradeCh {
		if err := conn.WriteJSON(trade); err != nil {
			log.Println("WS write failed:", err)
			return
		}
	}
}