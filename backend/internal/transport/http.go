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


func NewHTTPHandler(ob *orderbook.OrderBook) http.Handler {
	h := &HTTPHandler{ob: ob}

	mux := http.NewServeMux()
	mux.HandleFunc("/order", h.placeOrder)
	mux.HandleFunc("/orderbook", h.getOrderBook)

	return mux
}
