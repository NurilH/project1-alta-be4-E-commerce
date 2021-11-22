package databases

import (
	"fmt"
	"log"
	"project_altabe4_1/config"
	"project_altabe4_1/models"
)

func GetAllCart(id_user_token int) (interface{}, error) {
	cart := []models.Result{}
	where_clause := fmt.Sprintf("carts.users_id = %v", id_user_token)
	query := config.DB.Table("carts").Select("carts.id, carts.qty, carts.total_harga, carts.users_id, carts.product_id, products.nama, products.harga, products.kategori, products.deskripsi").Joins("join products on carts.product_id = products.id").Where(where_clause).Find(&cart)

	log.Println("result", where_clause)
	log.Println("result", cart)
	if query.Error != nil {
		return nil, query.Error
	}
	return cart, nil
}

func CreateCart(Cart *models.Cart) (interface{}, error) {

	if err := config.DB.Create(&Cart).Error; err != nil {
		return nil, err
	}

	return Cart.UsersID, nil
}

func GetHargaProduct(id int) (int, error) {
	product := models.Product{}
	err := config.DB.Find(&product, id)
	if err.Error != nil {
		return 0, err.Error
	}
	log.Println("harga", product.Harga)
	return product.Harga, nil
}

func UpdateCart(id int, Cart *models.Cart) {
	config.DB.Where("id = ?", id).Updates(&Cart)
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

func GetIDUserCart(id int) (uint, uint, error) {
	var cart models.Cart
	err := config.DB.Find(&cart, id)
	if err.Error != nil {
		return 0, 0, err.Error
	}
	return cart.UsersID, cart.ProductID, nil
}
