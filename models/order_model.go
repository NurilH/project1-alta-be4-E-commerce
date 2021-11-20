package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	StatusOrder bool `json:"statusorder" form:"statusorder"`
	TotalQty    int  `json:"totalqty" form:"totalqty"`
	CreditID    uint `json:"cartid" form:"cartid"`
}

type AddressRequest struct {
	Street string `json:"street" form:"street"`
	City   string `json:"city" form:"city"`
	State  string `json:"state" form:"state"`
	Zip    int    `json:"zip" form:"zip"`
}

type CreditCardRequest struct {
	Type   string `json:"type" form:"type"`
	Name   string `json:"name" form:"name"`
	Number int    `json:"number" form:"number"`
	Cvv    int    `json:"cvv" form:"cvv"`
	Month  int    `json:"month" form:"month"`
	Year   int    `json:"year" form:"year"`
}

type OrderRequest struct {
	CartId     []int             `json:"cart_id" form:"cart_id"`
	Address    AddressRequest    `json:"address" form:"address"`
	CreditCard CreditCardRequest `json:"credit_card" form:"credit_card"`
}
