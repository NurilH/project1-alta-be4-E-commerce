package controllers

import (
	"log"
	"net/http"
	"project_altabe4_1/models"

	"github.com/labstack/echo/v4"
)

func CreateOrderControllers(c echo.Context) error {

	order_req := models.OrderRequest{}

	c.Bind(&order_req)

	for _, v := range order_req.CartId {
		log.Println("id cart :", v)
	}

	log.Println("isi cart :", order_req.CartId)
	log.Println("isi address zip :", order_req.Address.Zip)
	log.Println("isi credit card :", order_req.CreditCard.Cvv)
	// Order := models.Order{}
	// c.Bind(&Order)
	// // log.Println(Order.ProductID)

	// logged := middlewares.ExtractTokenId(c)

	// // id_user_Order, _ := databases.GetIDUserProduct(int(Order.CartID))
	// // harga_product, _ := databases.GetHargaProduct(int(Order.CartID))

	// // Order.UsersID = uint(logged)
	// // Order.TotalHarga = Order.TotalQty * harga_product

	// if uint(logged) == id_user_Order {
	// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"code":    http.StatusBadRequest,
	// 		"message": "Access Forbidden",
	// 	})
	// }
	// _, e := databases.CreateOrder(&Order)
	// if e != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"code":    http.StatusBadRequest,
	// 		"message": "Bad Request",
	// 	})
	// }
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Successful Operation"})

}
