package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	StatusOrder    bool `json:"status_order" form:"status_order"`
	TotalQty       int  `json:"total_qty" form:"total_qty"`
	TotalHarga     int  `json:"total_harga" form:"total_harga"`
	CreditID       uint `json:"credit_id" form:"credit_id"`
	AddressRequest uint
	UsersID        uint
	DaftarOrder    []DaftarOrder
}

type DaftarOrder struct {
	gorm.Model
	OrderID uint
	CartID  uint
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
