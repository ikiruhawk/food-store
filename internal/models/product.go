package models

type Product struct {
	Id           int
	Name         string
	Category     Category
	Manufacturer Manufacturer
	Price        int
	PriceFloat   float64
	Amount       int
	Description  string
}

type Category struct {
	Id   int
	Name string
}

type Manufacturer struct {
	Id   int
	Name string
}
