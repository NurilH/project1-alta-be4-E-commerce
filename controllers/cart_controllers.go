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
		"message": "Successful Operation"})
}

func GetAllCartControllers(c echo.Context) error {
	logged := middlewares.ExtractTokenId(c)
	cart, err := databases.GetAllCart(logged)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Successful Operation",
		"data":    cart,
	})
}

func DeleteCartControllers(c echo.Context) error {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "False Param",
		})
	}
	id_user_cart, _, _ := databases.GetIDUserCart(conv_id)
	log.Println("id_user_cart", id_user_cart)
	logged := middlewares.ExtractTokenId(c)
	log.Println("idlogged", logged)
	if uint(logged) != id_user_cart {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Access Forbidden",
		})
	}
	databases.DeleteCart(conv_id)
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

	//mengecek user id nya sama dan ada pada tabel
	id_user_cart, id_product, _ := databases.GetIDUserCart(id)
	// log.Println("id_user_cart", id_user_cart)
	logged := middlewares.ExtractTokenId(c)
	// log.Println("idlogged", logged)
	if id_user_cart == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request",
		})
	} else if uint(logged) != id_user_cart {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Access Forbidden",
		})
	}

	//mengupdate total harga
	harga_product, _ := databases.GetHargaProduct(int(id_product))
	cart.TotalHarga = cart.Qty * harga_product

	//untuk mengupdate
	databases.UpdateCart(id, &cart)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Successful Operation",
	})
}
