package config

import (
	"os"

	"project_altabe4_1/models"

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
	DB.AutoMigrate(&models.Users{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Cart{})
	DB.AutoMigrate(&models.Credit{})
	DB.AutoMigrate(&models.Order{})
}

// ===============================================================//

func InitDBTest() {
	config_testing := os.Getenv("CONNECTION_DB_TESTING")

	var e error
	DB, e = gorm.Open(mysql.Open(config_testing), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrationTest()
}

func InitMigrationTest() {
	DB.Migrator().DropTable(&models.Users{})
	DB.AutoMigrate(&models.Users{})
}
