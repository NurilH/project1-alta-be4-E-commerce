package config

import (
	"os"

	"project_altabe4_1/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// inisialisasi database
func InitDB() {
	config := os.Getenv("CONNECTION_DB")

	var e error

	DB, e = gorm.Open(mysql.Open(config), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
}

// auto migrate -> untuk membuat tabel otomatis jika tabel tidak terdapat pada database
func InitMigrate() {
	DB.AutoMigrate(&models.Users{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Cart{})
	DB.AutoMigrate(&models.Credit{})
	DB.AutoMigrate(&models.Order{})
	DB.AutoMigrate(&models.AddressRequest{})
	DB.AutoMigrate(&models.DaftarOrder{})
}

// ===============================================================//

// inisialisasi database untuk untuk unit testing
func InitDBTest() {
	config_testing := os.Getenv("CONNECTION_DB_TESTING")

	var e error
	DB, e = gorm.Open(mysql.Open(config_testing), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrationTest()
}

// auto migrate -> untuk membuat tabel otomatis jika tabel tidak terdapat pada database
// drop table -> untuk menghapus tabel terlebih dahulu agar isi datanya dimulai dari tabel kosong
func InitMigrationTest() {
	DB.Migrator().DropTable(&models.Users{})
	DB.AutoMigrate(&models.Users{})
}
