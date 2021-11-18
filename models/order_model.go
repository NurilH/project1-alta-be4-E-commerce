package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	StatusOrder bool `json:"statusorder" form:"statusorder"`
	TotalQty    int  `json:"totalqty" form:"totalqty"`
}
