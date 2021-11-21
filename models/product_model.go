package models

import "gorm.io/gorm"

// struktur data product
type Product struct {
	gorm.Model
	Nama      string `json:"nama" form:"nama"`
	Harga     int    `json:"harga" form:"harga"`
	Kategori  string `json:"kategori" form:"kategori"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
	UsersID   uint
	Cart      []Cart
}
