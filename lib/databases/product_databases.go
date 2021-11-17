package databases

import (
	"project_altabe4_1/config"
	"project_altabe4_1/models"
)

func GetAllProduct() (interface{}, error) {
	products := []models.Product{}
	err := config.DB.Find(products)
	if err.Error != nil {
		return nil, err.Error
	}
	return products, nil
}

func CreateProduct(product *models.Product) (interface{}, error) {
	if err := config.DB.Create(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func UpdateProduct(id int, products *models.Product) (interface{}, error) {
	if err := config.DB.Where("id = ?", id).Updates(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
