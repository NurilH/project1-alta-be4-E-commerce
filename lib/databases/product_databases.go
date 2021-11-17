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
