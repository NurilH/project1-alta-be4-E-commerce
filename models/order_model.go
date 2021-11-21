package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	StatusOrder bool `json:"status_order" form:"status_order"`
	TotalQty    int  `json:"total_qty" form:"total_qty"`
	CreditID    uint `json:"credit_id" form:"credit_id"`
	DetailID    uint
}

type DaftarOrder struct {
	gorm.Model
	OrderID          uint
	AddressRequestID uint
	DetailCartId     int
}

type AddressRequest struct {
	gorm.Model
	Street string `json:"street" form:"street"`
	City   string `json:"city" form:"city"`
	State  string `json:"state" form:"state"`
	Zip    int    `json:"zip" form:"zip"`
}

type OrderRequest struct {
	DetailCartId []int          `json:"detail_cart_id" form:"detail_cart_id"`
	Order        Order          `json:"order" form:"order"`
	Address      AddressRequest `json:"address" form:"address"`
}

// type CreditCardRequest struct {
// 	Type   string `json:"type" form:"type"`
// 	Name   string `json:"name" form:"name"`
// 	Number int    `json:"number" form:"number"`
// 	Cvv    int    `json:"cvv" form:"cvv"`
// 	Month  int    `json:"month" form:"month"`
// 	Year   int    `json:"year" form:"year"`
// }

// type OrderRequest struct {
// 	gorm.Model
// 	CartId  int            `json:"cart_id" form:"cart_id"`
// 	Address AddressRequest `json:"address" form:"address"`
// 	CreditCard CreditCardRequest `json:"credit_card" form:"credit_card"`
// }
