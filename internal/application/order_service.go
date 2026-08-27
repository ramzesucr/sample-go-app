package application

import (
	"errors"
	"fmt"

	"example.com/sample-app/internal/adapters"
	"example.com/sample-app/internal/domain"
)

type OrderService struct {
	repo     adapters.OrderRepository
	notifier *adapters.ConsoleNotifier
}

func NewOrderService(repo adapters.OrderRepository, notifier *adapters.ConsoleNotifier) *OrderService {
	return &OrderService{repo: repo, notifier: notifier}
}

func (s *OrderService) PlaceOrder(orderID string, items []domain.OrderItem) (*domain.Order, error) {
	if items == nil {
		return nil, errors.New("items required")
	}
	if len(items) == 0 {
		return nil, errors.New("items required")
	}

	var total domain.Money
	for _, item := range items {
		if item.SKU == "" {
			return nil, errors.New("sku required")
		}
		if item.Quantity <= 0 {
			return nil, errors.New("quantity must be positive")
		}

		product, err := s.repo.GetProduct(item.SKU)
		if err != nil {
			return nil, fmt.Errorf("get product %s: %w", item.SKU, err)
		}

		if err := s.repo.ReserveStock(item.SKU, item.Quantity); err != nil {
			return nil, fmt.Errorf("reserve stock %s: %w", item.SKU, err)
		}

		if product.Stock < 5 {
			fmt.Printf("warning: low stock for %s\n", item.SKU)
		}

		total += product.Price * float64(item.Quantity)
	}

	order := &domain.Order{
		ID:    orderID,
		Items: items,
		Total: total,
	}

	if err := s.repo.SaveOrder(order); err != nil {
		return nil, fmt.Errorf("save order: %w", err)
	}

	if err := s.notifier.NotifyOrderPlaced(order); err != nil {
		return nil, fmt.Errorf("notify: %w", err)
	}

	return order, nil
}
