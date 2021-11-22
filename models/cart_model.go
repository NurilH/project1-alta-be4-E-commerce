package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Qty        int  `json:"qty" form:"qty"`
	TotalHarga int  `json:"totalharga" form:"totalharga"`
	ProductID  uint `json:"productid" form:"productid"`
	UsersID    uint
}

type Result struct {
	ID         uint   `json:"id"`
	Qty        int    `json:"qty"`
	TotalHarga int    `json:"total_harga"`
	UsersID    uint   `json:"users_id"`
	ProductID  uint   `json:"product_id"`
	Nama       string `json:"nama"`
	Harga      int    `json:"harga"`
	Kategori   string `json:"kategori"`
	Deskripsi  string `json:"deskripsi"`
}
