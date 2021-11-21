package controllers

import (
	"log"
	"net/http"
	"project_altabe4_1/lib/databases"
	"project_altabe4_1/models"

	"github.com/labstack/echo/v4"
)

func CreateOrderControllers(c echo.Context) error {
	order_req := models.OrderRequest{}
	c.Bind(&order_req)

	// var d interface{}
	// var er error
	for _, v := range order_req.DetailCartId {
		log.Println("id detail detail", v)
		order := models.Order{}
		order.DetailId = v
		order.TotalQty = order_req.Order.TotalQty
		order.CreditID = order_req.Order.CreditID
		order.StatusOrder = order_req.Order.StatusOrder
		databases.CreateOrderDet(&order)
	}

	order_detail, er := databases.CreateOrder(&order_req)
	if er != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request",
		})
	}

	log.Println("isi cart :", order_req.DetailCartId)
	log.Println("isi city :", order_req.Address.City)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Successful Operation",
		"address": order_detail,
	})

}
