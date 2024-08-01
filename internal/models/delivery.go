package models

import "time"

type Delivery struct {
	Id    int       `json:"id"`
	Store Store     `json:"store"`
	Date  time.Time `json:"date"`
}

type DeliveryProducts struct {
	Delivery Delivery `json:"delivery"`
	Product  Product  `json:"product"`
	Amount   int      `json:"amount"`
}
