package databases

import (
	"log"
	"project_altabe4_1/config"
	"project_altabe4_1/models"
)

func GetAllCart() (interface{}, error) {
	type result struct {
		CartID           uint
		Qty              int
		TotalHarga       int
		UsersID          uint
		ProductID        uint
		NamaProduct      string
		HargaProduct     int
		DeskripsiProduct string
	}
	carts := models.Cart{}
	cart := config.DB.Joins("").Find(&carts)
	// cart := config.DB.Model(&models.Cart{}).Select("cart.id, cart.qty, cart.total_harga, cart.users_id, product.id, product.nama, product.harga, product.deskripsi").Joins("inner join product on cart.product_id = product.id").Scan(result)
	log.Println("result", result{})
	log.Println("result", cart)
	if cart.Error != nil {
		return nil, cart.Error
	}
	return carts, nil
}
