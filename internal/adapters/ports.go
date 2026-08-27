package adapters

import "example.com/sample-app/internal/domain"

// OrderRepository persists orders and reserves stock.
type OrderRepository interface {
	GetProduct(sku string) (*domain.Product, error)
	ReserveStock(sku string, qty int) error
	SaveOrder(order *domain.Order) error
}

// Notifier sends order confirmations.
type Notifier interface {
	NotifyOrderPlaced(order *domain.Order) error
}
