package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Qty        int  `json:"qty" form:"qty"`
	TotalHarga int  `json:"totalharga" form:"totalharga"`
	ProductID  uint `json:"productid" form:"productid"`
	UsersID    uint
	Order      []Order
}
