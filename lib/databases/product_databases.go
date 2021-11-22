package databases

import (
	"project_altabe4_1/config"
	"project_altabe4_1/models"
)

func CreateProduct(product *models.Product) (interface{}, error) {
	if err := config.DB.Create(&product).Error; err != nil {
		return nil, err
	}
	return product.UsersID, nil
}

func GetAllProduct() (interface{}, error) {
	products := []models.Product{}
	res := []models.Get_Products{}
	err := config.DB.Model(&products).Find(&res)
	if err.Error != nil {
		return nil, err.Error
	}
	return res, nil
}

func GetProductById(id int) (interface{}, error) {
	product := models.Product{}
	res := models.Get_Products{}
	err := config.DB.Model(&product).Find(&res, id)
	rows_affected := config.DB.Find(&product, id).RowsAffected
	if err.Error != nil || rows_affected < 1 {
		return nil, err.Error
	}
	return res, nil
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
	if err != nil || check_product > 0 {
		return nil, err
	}
	return product.UsersID, nil
}

func UpdateProduct(id int, products *models.Product) (interface{}, error) {
	if err := config.DB.Where("id = ?", id).Updates(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
