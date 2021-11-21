package controllers

import (
	"net/http"
	"project_altabe4_1/lib/databases"
	"project_altabe4_1/middlewares"
	"project_altabe4_1/models"
	"project_altabe4_1/response"
	"strconv"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// controller untuk membuat cart
func CreateCartControllers(c echo.Context) error {

	Cart := models.Cart{}
	c.Bind(&Cart)
	v := validator.New()
	e := v.Var(Cart.Qty, "required,gt=0")
	if e == nil {
		logged := middlewares.ExtractTokenId(c)

		id_user_cart, _ := databases.GetIDUserProduct(int(Cart.ProductID))
		harga_product, _ := databases.GetHargaProduct(int(Cart.ProductID))

		Cart.UsersID = uint(logged)
		Cart.TotalHarga = Cart.Qty * harga_product

		if uint(logged) == id_user_cart {
			return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
		}
		_, e = databases.CreateCart(&Cart)
	}
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

// controller untuk mendapatkan semua cart by id user
func GetAllCartControllers(c echo.Context) error {
	logged := middlewares.ExtractTokenId(c)
	cart, err := databases.GetAllCart(logged)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(cart))
}

// controller untuk menghapus cart by id
func DeleteCartControllers(c echo.Context) error {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	id_user_cart, _, _ := databases.GetIDUserCart(conv_id)
	logged := middlewares.ExtractTokenId(c)
	if uint(logged) != id_user_cart {
		return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
	}
	databases.DeleteCart(conv_id)
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

func UpdateCartControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}

	cart := models.Cart{}
	c.Bind(&cart)
	v := validator.New()
	e := v.Var(cart.Qty, "required,gt=0")
	if e == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	// mengecek user id nya sama dan ada pada tabel
	id_user_cart, id_product, _ := databases.GetIDUserCart(id)
	logged := middlewares.ExtractTokenId(c)
	if id_user_cart == 0 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	} else if uint(logged) != id_user_cart {
		return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
	}

	// mengupdate total harga
	harga_product, _ := databases.GetHargaProduct(int(id_product))
	cart.TotalHarga = cart.Qty * harga_product

	// untuk mengupdate
	databases.UpdateCart(id, &cart)

	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}
