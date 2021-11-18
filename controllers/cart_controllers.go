package controllers

import (
	"net/http"
	"project_altabe4_1/lib/databases"

	"github.com/labstack/echo/v4"
)

func GetAllCartControllers(c echo.Context) error {
	cart, err := databases.GetAllCart()
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
