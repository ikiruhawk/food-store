package models

type Product struct {
	Id           int          `json:"id"`
	Name         string       `json:"name"`
	Category     Category     `json:"category"`
	Manufacturer Manufacturer `json:"manufacturer"`
	Price        int          `json:"price"`
	PriceFloat   float64      `json:"priceFloat"`
	Amount       int          `json:"amount"`
	Description  string       `json:"description"`
}

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Manufacturer struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
