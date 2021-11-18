package databases

import (
	"project_altabe4_1/config"
	"project_altabe4_1/models"
)

func CreateOrder(Order *models.Order) (interface{}, error) {

	if err := config.DB.Create(&Order).Error; err != nil {
		return nil, err
	}

	return Order, nil
}
