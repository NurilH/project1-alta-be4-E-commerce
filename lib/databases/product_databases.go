package databases

import (
	"project_altabe4_1/config"
	"project_altabe4_1/models"
)

func CreateProduct(user *models.Product) (interface{}, error) {
	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
