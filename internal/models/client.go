package models

type Client struct {
	Id          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	Login       string `json:"login"`
	Password    string `json:"password"`
}

type ClientAddress struct {
	Id      int    `json:"id"`
	Client  Client `json:"client"`
	Address string `json:"address"`
}
