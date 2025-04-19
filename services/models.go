package services

// Package services defines data models used in the application.
// These models represent the structure of assets and holdings in the investment dashboard.
// They are used to map database records to Go structs, making it easier to work with data in the application.

// Asset represents a financial asset with its details.
// It includes fields for symbol, name, quantity, value, 24-hour change, and allocation.
// This struct is used to represent individual assets in a user's portfolio.

type Asset struct {
	Symbol     string  `json:"symbol"`
	Name       string  `json:"name"`
	Quantity   float64 `json:"quantity"`
	Value      float64 `json:"value"`
	Change24h  float64 `json:"change24h"`
	Allocation float64 `json:"allocation"`
}

// Holding represents a user's holding in a portfolio.
// It includes fields for ID, portfolio ID, symbol, name, quantity, and value.
// This struct is used to represent the holdings of a user, which are part of their investment portfolio.

type Holding struct {
	ID          string  `json:"id"`
	PortfolioID string  `json:"portfolio_id"`
	Symbol      string  `json:"symbol"`
	Name        string  `json:"name"`
	Quantity    float64 `json:"quantity"`
	Value       float64 `json:"value"`
}
