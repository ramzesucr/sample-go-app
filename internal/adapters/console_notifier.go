package adapters

import (
	"fmt"

	"example.com/sample-app/internal/domain"
)

type ConsoleNotifier struct{}

func NewConsoleNotifier() *ConsoleNotifier {
	return &ConsoleNotifier{}
}

func (n *ConsoleNotifier) NotifyOrderPlaced(order *domain.Order) error {
	fmt.Printf("order %s placed, total: %.2f\n", order.ID, order.Total)
	return nil
}
