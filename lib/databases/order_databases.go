package databases

import (
	"project_altabe4_1/config"
	"project_altabe4_1/models"
)

func CreateOrder(Address *models.OrderRequest) (interface{}, error) {
	config.DB.Create(&Address.Order)
	if err := config.DB.Create(&Address.Address).Error; err != nil {
		return nil, err
	}

	return Address, nil
}
