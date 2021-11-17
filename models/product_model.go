package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Nama      string `json:"nama" form:"nama"`
	Harga     string `json:"harga" form:"harga"`
	Kategori  string `json:"kategori" form:"kategori"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
	UsersID   int
}
