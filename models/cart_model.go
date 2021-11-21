package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Qty         int  `json:"qty" form:"qty"`
	TotalHarga  int  `json:"total_harga" form:"total_harga"`
	ProductID   uint `json:"product_id" form:"product_id"`
	UsersID     uint
	DaftarOrder []DaftarOrder
}
