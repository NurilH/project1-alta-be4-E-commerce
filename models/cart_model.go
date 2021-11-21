package models

import "gorm.io/gorm"

// struktur data cart
type Cart struct {
	gorm.Model
	Qty        int  `json:"qty" form:"qty"`
	TotalHarga int  `json:"totalharga" form:"totalharga"`
	ProductID  uint `json:"productid" form:"productid"`
	UsersID    uint
}
