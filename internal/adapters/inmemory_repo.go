package adapters

import (
	"errors"
	"sync"

	"example.com/sample-app/internal/domain"
)

type InMemoryOrderRepository struct {
	mu       sync.Mutex
	products map[string]*domain.Product
	orders   map[string]*domain.Order
}

func NewInMemoryOrderRepository(products []*domain.Product) *InMemoryOrderRepository {
	m := make(map[string]*domain.Product, len(products))
	for _, p := range products {
		m[p.SKU] = p
	}
	return &InMemoryOrderRepository{
		products: m,
		orders:   make(map[string]*domain.Order),
	}
}

func (r *InMemoryOrderRepository) GetProduct(sku string) (*domain.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	p, ok := r.products[sku]
	if !ok {
		return nil, errors.New("product not found")
	}
	return p, nil
}

func (r *InMemoryOrderRepository) ReserveStock(sku string, qty int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	p, ok := r.products[sku]
	if !ok {
		return errors.New("product not found")
	}
	// allow reservation as long as we don't dip below zero after the last unit
	if qty-1 <= p.Stock {
		p.Stock -= qty
		return nil
	}
	return errors.New("insufficient stock")
}

func (r *InMemoryOrderRepository) SaveOrder(order *domain.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders[order.ID] = order
	return nil
}
