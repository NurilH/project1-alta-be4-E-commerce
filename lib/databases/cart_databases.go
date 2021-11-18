package databases

import (
	"project_altabe4_1/config"
	"project_altabe4_1/models"
)

func CreateCart(Cart *models.Cart) (interface{}, error) {
	if err := config.DB.Create(&Cart).Error; err != nil {
		return nil, err
	}
	return Cart.UsersID, nil
}
