package models

type Client struct {
	Id          int
	FirstName   string
	LastName    string
	PhoneNumber string
	Login       string
	Password    string
}

type ClientAddress struct {
	Id      int
	Client  Client
	Address string
}
