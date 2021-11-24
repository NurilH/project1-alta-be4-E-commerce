package databases

import (
	"log"
	"project_altabe4_1/config"
	"project_altabe4_1/models"
)

type ResponseOrder struct {
	CreditID    uint
	AddressID   uint
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

type CartOrder struct {
	ID    uint   `json:"id"`
	Nama  string `json:"nama"`
	Qty   int    `json:"qty"`
	Harga int    `json:"harga"`
}

func CreateOrder(Order_a *models.OrderRequest) (interface{}, error) {

	if err := config.DB.Create(&Order_a.Order).Error; err != nil {
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
			Order_a.Order.CreditID,
			Order_a.Order.AddressRequest,
			Order_a.Order.StatusOrder,
		},
	}, nil
}

func CreateAddress(address *models.AddressRequest) {
	config.DB.Create(&address)
}

func CreateOrderDet(Order *models.DaftarOrder) (interface{}, error) {
	if err := config.DB.Create(&Order).Error; err != nil {
		return nil, err
	}
	return Order, nil
}

func GetHargaQtyCart(id int) (int, int, error) {
	cart := models.Cart{}
	err := config.DB.Find(&cart, id)
	if err.Error != nil {
		return 0, 0, err.Error
	}
	log.Println("harga", cart.TotalHarga)
	return cart.TotalHarga, cart.Qty, nil
}

func GetOrder(id int) (interface{}, interface{}, error) {
	order := models.DaftarOrder{}
	type ord struct {
		ID         uint
		Userid     uint
		TotalQty   int
		TotalHarga int
	}

	order_a := models.Order{}

	config.DB.Where("credit_id = ?", 3).Find(&order_a)
	log.Println("id dari order", id)
	config.DB.Where("order_id <> ?", order_a.ID).Find(&order)
	// cart_order := CartOrder{}

	// query := config.DB.Table("daftar_orders").Select("products.id, products.nama, carts.qty, carts.total_harga").Joins("join carts on daftar_orders.detail_cart_id = carts.id ").Joins("join products on carts.product_id = products.id").Where("daftar_orders.order_id=2").Find(&cart_order)

	// log.Println("query ", query)
	// log.Println("cart ", cart_order)
	// err := config.DB.Find(&order, id)

	// if query.Error != nil {
	// 	return nil, query.Error
	// }

	return order, ord{order_a.ID, order_a.UsersID, order_a.TotalQty, order_a.TotalHarga}, nil
}
