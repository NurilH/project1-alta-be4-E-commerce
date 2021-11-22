package models

import "gorm.io/gorm"

type Credit struct {
	gorm.Model
	Typ     string `json:"typ" form:"typ"`
	Nama    string `json:"nama" form:"nama"`
	Nomor   string `json:"nomor" form:"nomor"`
	Cvv     int    `json:"cvv" form:"cvv"`
	Bulan   int    `json:"bulan" form:"bulan"`
	Tahun   int    `json:"tahun" form:"tahun"`
	UsersID uint
	Order   []Order
}
