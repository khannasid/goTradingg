/*
If asked:
“Where does the goroutine for the WebSocket client come from?”

You can confidently say:
“In Go’s net/http server, each request handler already runs in its own goroutine. 
The WebSocket connection lives for the lifetime of that handler, 
so the blocking read loop naturally runs inside that goroutine without needing 
to spawn a separate one.”
*/

package transport

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"

	"goTradingg/internal/models"
	"goTradingg/internal/orderbook"
)

type HTTPHandler struct {
	ob *orderbook.OrderBook
	broadcaster *TradeBroadcaster
}

type PlaceOrderRequest struct {
	Side     models.OrderSide `json:"side"`
	Price    float64          `json:"price"`
	Quantity int              `json:"quantity"`
}

func (h *HTTPHandler) placeOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req PlaceOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	order := models.Order{
		ID:        uuid.New().String(),
		Symbol:    h.ob.Symbol,
		Side:      req.Side,
		Price:     req.Price,
		Quantity:  req.Quantity,
		Timestamp: time.Now(),
	}

	h.ob.AddOrder(order)
	trades := h.ob.Match()

	for _, trade := range trades{
		h.broadcaster.Broadcast(trade)
	}

	response := map[string]interface{}{
		"order_id": order.ID,
		"trades":   trades,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *HTTPHandler) getOrderBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	buys, sells := h.ob.Snapshot()
	json.NewEncoder(w).Encode(map[string]interface{}{
		"buy_orders":  buys,
		"sell_orders": sells,
	})
}


func NewHTTPHandler(ob *orderbook.OrderBook, b *TradeBroadcaster) http.Handler {
	h := &HTTPHandler{ob: ob, broadcaster: b}

	mux := http.NewServeMux()
	mux.HandleFunc("/order", h.placeOrder)
	mux.HandleFunc("/orderbook", h.getOrderBook)
	mux.HandleFunc("/ws/trades", h.tradeWS)

	return withCORS(mux)
}
