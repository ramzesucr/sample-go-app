package domain

type OrderItem struct {
	SKU      string
	Quantity int
}

type Order struct {
	ID    string
	Items []OrderItem
	Total Money
}
