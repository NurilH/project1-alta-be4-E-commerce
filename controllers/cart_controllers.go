package controllers

import (
	"log"
	"net/http"
	"project_altabe4_1/lib/databases"
	"project_altabe4_1/middlewares"
	"project_altabe4_1/models"

	"github.com/labstack/echo/v4"
)

func CreateCartControllers(c echo.Context) error {
	Cart := models.Cart{}
	c.Bind(&Cart)
	log.Println(Cart.ProductID)

	logged := middlewares.ExtractTokenId(c)

	id_user_cart, _ := databases.GetIDUserProduct(int(Cart.ProductID))
	Cart.UsersID = uint(logged)
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
