package controllers

import (
	"net/http"
	"project_altabe4_1/lib/databases"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetProductByIdControllers(c echo.Context) error {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "False Param",
		})
	}
	product, e := databases.GetProductById(conv_id)
	if e != nil || product == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Successful Operation",
		"data":    product,
	})
}
