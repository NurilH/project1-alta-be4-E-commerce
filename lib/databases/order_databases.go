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

func CreateOrder(Order_a *models.OrderRequest) (interface{}, error) {
	config.DB.Create(&Order_a.Order)
	if err := config.DB.Create(&Order_a.Address).Error; err != nil {
		return nil, err
	}
	return OrderDetailRequest{Order_a.DetailCartId,
		AddressRequest{
			Order_a.Address.Street,
			Order_a.Address.City,
			Order_a.Address.State,
			Order_a.Address.Zip,
		},
		ResponseOrder{
			Order_a.Order.TotalQty,
			Order_a.Order.CreditID,
			Order_a.Order.StatusOrder,
		},
	}, nil
}

func CreateOrderDet(Order *models.DaftarOrder) (interface{}, error) {
	if err := config.DB.Create(&Order).Error; err != nil {
		return nil, err
	}

	return Order, nil
}
