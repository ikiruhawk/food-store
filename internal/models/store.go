package models

type Store struct {
	Id      int
	Address string
}

type StoreProducts struct {
	Store   Store
	Product Product
	Amount  int
}
