package models

import "gorm.io/gorm"

type Detail struct {
	gorm.Model
	CartID  uint `json:"cart_id" form:"cart_id"`
	OrderID uint `json:"order_id" form:"order_id"`
}
