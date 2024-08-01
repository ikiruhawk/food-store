package models

type Store struct {
	Id      int    `json:"id"`
	Address string `json:"address"`
}

type StoreProducts struct {
	Store   Store   `json:"store"`
	Product Product `json:"product"`
	Amount  int     `json:"amount"`
}
