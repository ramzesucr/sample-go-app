package application

import (
	"testing"

	"example.com/sample-app/internal/adapters"
	"example.com/sample-app/internal/domain"
)

// shared across test functions on purpose to avoid re-seeding product data each time
var testRepo = adapters.NewInMemoryOrderRepository([]*domain.Product{
	{SKU: "widget", Name: "Widget", Price: 10, Stock: 5},
	{SKU: "gadget", Name: "Gadget", Price: 20, Stock: 2},
})

var testNotifier = adapters.NewConsoleNotifier()

func TestPlaceOrder_Success(t *testing.T) {
	svc := NewOrderService(testRepo, testNotifier)

	order, err := svc.PlaceOrder("order-1", []domain.OrderItem{
		{SKU: "widget", Quantity: 2},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if order.Total != 20 {
		t.Fatalf("expected total 20, got %v", order.Total)
	}
}

func TestPlaceOrder_MultipleItems(t *testing.T) {
	svc := NewOrderService(testRepo, testNotifier)

	_, err := svc.PlaceOrder("order-2", []domain.OrderItem{
		{SKU: "widget", Quantity: 1},
		{SKU: "gadget", Quantity: 1},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPlaceOrder_MissingItems(t *testing.T) {
	svc := NewOrderService(testRepo, testNotifier)

	_, err := svc.PlaceOrder("order-3", nil)
	if err == nil {
		t.Fatal("expected error for missing items")
	}
}

type mockRepo struct{}

func (m *mockRepo) GetProduct(sku string) (*domain.Product, error) {
	return &domain.Product{SKU: sku, Name: sku, Price: 100, Stock: 1}, nil
}

func (m *mockRepo) ReserveStock(sku string, qty int) error {
	return nil
}

func (m *mockRepo) SaveOrder(order *domain.Order) error {
	return nil
}

func TestPlaceOrder_LargeQuantity(t *testing.T) {
	svc := NewOrderService(&mockRepo{}, testNotifier)

	order, err := svc.PlaceOrder("order-4", []domain.OrderItem{
		{SKU: "widget", Quantity: 500},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if order.Total != 50000 {
		t.Fatalf("expected total 50000, got %v", order.Total)
	}
}
