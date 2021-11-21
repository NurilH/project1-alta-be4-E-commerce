package databases

import (
	"project_altabe4_1/config"
	"project_altabe4_1/models"
)

type ResponseOrder struct {
	TotalQty    int
	CreditID    uint
	StatusOrder bool
}
type AddressRequest struct {
	Street string
	City   string
	State  string
	Zip    int
}
type OrderDetailRequest struct {
	DetailCartId []int
	Address      AddressRequest
	Order        ResponseOrder
}

func CreateOrder(Address *models.OrderRequest) (interface{}, error) {
	if err := config.DB.Create(&Address.Address).Error; err != nil {
		return nil, err
	}
	return OrderDetailRequest{Address.DetailCartId,
		AddressRequest{
			Address.Address.Street,
			Address.Address.City,
			Address.Address.State,
			Address.Address.Zip,
		},
		ResponseOrder{
			Address.Order.TotalQty,
			Address.Order.CreditID,
			Address.Order.StatusOrder,
		},
	}, nil
}

func CreateOrderDet(Order *models.Order) (interface{}, error) {
	if err := config.DB.Create(&Order).Error; err != nil {
		return nil, err
	}

	return Order, nil
}
