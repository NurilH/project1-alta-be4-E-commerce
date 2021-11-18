package models

import (
	"time"

	"gorm.io/gorm"
)

type Credit struct {
	gorm.Model
	Typ        string    `json:"typ" form:"typ"`
	CardNumber int       `json:"cardnumber" form:"cardnumber"`
	Cvv        int       `json:"cvv" form:"cvv"`
	MasaAktif  time.Time `json:"masaaktif" form:"masaaktif"`
	UsersID    uint
}
