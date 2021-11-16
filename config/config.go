package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	config := os.Getenv("CONNECTION_DB")

	var e error

	DB, e = gorm.Open(mysql.Open(config), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
}

func InitMigrate() {

}
