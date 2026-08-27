package main

import (
	"log"
	"net/http"

	"example.com/sample-app/internal/adapters"
	"example.com/sample-app/internal/application"
	"example.com/sample-app/internal/domain"
)

func main() {
	products := []*domain.Product{
		{SKU: "widget", Name: "Widget", Price: 9.99, Stock: 10},
		{SKU: "gadget", Name: "Gadget", Price: 19.99, Stock: 3},
	}

	repo := adapters.NewInMemoryOrderRepository(products)
	notifier := adapters.NewConsoleNotifier()
	service := application.NewOrderService(repo, notifier)
	handler := adapters.NewOrderHandler(service)

	http.HandleFunc("/orders", handler.PlaceOrder)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
