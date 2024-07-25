package models

import "time"

type Delivery struct {
	Id    int
	Store Store
	Date  time.Time
}

type DeliveryProducts struct {
	Delivery Delivery
	Product  Product
	Amount   int
}
