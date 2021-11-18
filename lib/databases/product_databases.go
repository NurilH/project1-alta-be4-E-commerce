package databases

import (
	"project_altabe4_1/config"
	"project_altabe4_1/models"
)

// a
// a

/*

































 */

func CreateProduct(product *models.Product) (interface{}, error) {
	if err := config.DB.Create(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func GetAllProduct() (interface{}, error) {
	products := []models.Product{}
	err := config.DB.Find(products)
	if err.Error != nil {
		return nil, err.Error
	}
	return products, nil
}

func GetProductById(id int) (interface{}, error) {
	product := models.Product{}
	type get_product struct {
		ID        uint
		Nama      string
		Harga     int
		Deskripsi string
	}
	err := config.DB.Find(&product, id)
	rows_affected := config.DB.Find(&product, id).RowsAffected
	if err.Error != nil || rows_affected < 1 {
		return nil, err.Error
	}
	return get_product{product.ID, product.Nama, product.Harga, product.Deskripsi}, nil
}

func GetIDUserProduct(id int) (uint, error) {
	var product models.Product
	err := config.DB.Find(&product, id)
	if err.Error != nil {
		return 0, err.Error
	}
	return product.UsersID, nil
}

func DeleteProduct(id int) (interface{}, error) {
	var product models.Product
	check_product := config.DB.Find(&product, id).RowsAffected

	err := config.DB.Delete(&product).Error
	// rows_affected := config.DB.Delete(&product).RowsAffected
	// log.Println("rows affected", rows_affected)
	if err != nil || check_product > 0 {
		return nil, err
	}
	return product.UsersID, nil
}
