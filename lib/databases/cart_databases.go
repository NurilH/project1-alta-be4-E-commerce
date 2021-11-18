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
	// cart := config.DB.Joins("").Find(&carts)
	cart := config.DB.Model(&models.Cart{}).Select("carts.id, carts.qty, carts.total_harga, carts.users_id, products.id, products.nama, products.harga, products.deskripsi").Joins("left join products on carts.product_id = products.id").Scan(&result{})
	// log.Println("result", carts)
	log.Println("result", cart)
	if cart.Error != nil {
		return nil, cart.Error
	}
	return cart, nil
}

func CreateCart(Cart *models.Cart) (interface{}, error) {
	if err := config.DB.Create(&Cart).Error; err != nil {
		return nil, err
	}
	return Cart.UsersID, nil
}

func DeleteCart(id int) (interface{}, error) {
	var cart models.Cart
	check_cart := config.DB.Find(&cart, id).RowsAffected
	err := config.DB.Delete(&cart).Error
	if err != nil || check_cart > 0 {
		return nil, err
	}
	return cart.UsersID, nil
}

func GetIDUserCart(id int) (uint, error) {
	var cart models.Cart
	err := config.DB.Find(&cart, id)
	if err.Error != nil {
		return 0, err.Error
	}
	return cart.UsersID, nil
}
