package models

import "time"

type Order struct {
	Id            int           `json:"id"`
	ClientAddress ClientAddress `json:"clientAddress"`
	Date          time.Time     `json:"date"`
	Price         int           `json:"price"`
}

type OrderProducts struct {
	Id      int     `json:"id"`
	Order   Order   `json:"order"`
	Product Product `json:"product"`
	Amount  int     `json:"amount"`
}
