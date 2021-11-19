package databases

import (
	"fmt"
	"log"
	"project_altabe4_1/config"
	"project_altabe4_1/models"
)

func GetAllCart(id_user_token int) (interface{}, error) {
	type result struct {
		ID         uint   `json:"id"`
		Qty        int    `json:"qty"`
		TotalHarga int    `json:"total_harga"`
		UsersID    uint   `json:"users_id"`
		ProductID  uint   `json:"product_id"`
		Nama       string `json:"nama"`
		Harga      int    `json:"harga"`
		Kategori   string `json:"kategori"`
		Deskripsi  string `json:"deskripsi"`
	}
	// cart := config.DB.Joins("").Find(&carts)
	// cart := config.DB.Model(&models.Cart{}).Select("carts.id, carts.qty, carts.total_harga, carts.users_id, products.id, products.nama, products.harga, products.deskripsi").Joins("left join products on carts.product_id = products.id").Scan(&result{})
	// log.Println("result", carts)
	cart := []result{}
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
