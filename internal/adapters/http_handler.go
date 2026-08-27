package adapters

import (
	"encoding/json"
	"net/http"

	"example.com/sample-app/internal/domain"
)

type OrderPlacer interface {
	PlaceOrder(orderID string, items []domain.OrderItem) (*domain.Order, error)
}

type OrderHandler struct {
	service OrderPlacer
}

func NewOrderHandler(service OrderPlacer) *OrderHandler {
	return &OrderHandler{service: service}
}

type placeOrderRequest struct {
	OrderID string             `json:"order_id"`
	Items   []domain.OrderItem `json:"items"`
}

func (h *OrderHandler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var req placeOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order, err := h.service.PlaceOrder(req.OrderID, req.Items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
