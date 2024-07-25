package models

import "time"

type Order struct {
	Id            int
	ClientAddress ClientAddress
	Date          time.Time
	Price         int
}

type OrderProducts struct {
	Id      int
	Order   Order
	Product Product
	Amount  int
}
