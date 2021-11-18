package controllers

import (
	"log"
	"net/http"
	"project_altabe4_1/lib/databases"
	"project_altabe4_1/middlewares"
	"project_altabe4_1/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateCartControllers(c echo.Context) error {

	Cart := models.Cart{}
	c.Bind(&Cart)
	log.Println(Cart.ProductID)

	logged := middlewares.ExtractTokenId(c)

	id_user_cart, _ := databases.GetIDUserProduct(int(Cart.ProductID))
	harga_product, _ := databases.GetHargaProduct(int(Cart.ProductID))

	Cart.UsersID = uint(logged)
	Cart.TotalHarga = Cart.Qty * harga_product

	if uint(logged) == id_user_cart {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Access Forbidden",
		})
	}
	_, e := databases.CreateCart(&Cart)
	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Successful Operation",
	})
}

func UpdateCartControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "False Param",
		})
	}

	cart := models.Cart{}
	c.Bind(&cart)
	// logged := middlewares.ExtractTokenId(c)
	// cart.UsersID = uint(logged)

	_, er := databases.UpdateCart(id, &cart)
	if er != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Successful Operation",
	})
}
