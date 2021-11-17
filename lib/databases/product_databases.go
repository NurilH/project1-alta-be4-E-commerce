package databases

import (
	"project_altabe4_1/config"
	"project_altabe4_1/models"
)

func CreateProduct(product *models.Product) (interface{}, error) {
	if err := config.DB.Create(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
