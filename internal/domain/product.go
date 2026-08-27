package domain

// Money represents a price in the store's default currency.
type Money = float64

type Product struct {
	SKU   string
	Name  string
	Price Money
	Stock int
}
