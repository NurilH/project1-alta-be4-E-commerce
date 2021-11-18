package databases

import (
	"log"
	"project_altabe4_1/config"
	"project_altabe4_1/models"
)

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

func UpdateCart(id int, Cart *models.Cart) (interface{}, error) {
	RowsAffected := config.DB.Where("id = ?", id).Updates(&Cart).RowsAffected
	err := config.DB.Where("id = ?", id).Updates(&Cart).Error
	if RowsAffected == 0 || err != nil {
		return nil, err
	}
	config.DB.First(&Cart, id)
	log.Println(RowsAffected)
	return Cart, nil
}
