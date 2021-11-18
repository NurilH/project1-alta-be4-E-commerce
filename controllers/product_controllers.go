package controllers

import (
	"log"
	"net/http"
	"project_altabe4_1/lib/databases"
	"project_altabe4_1/middlewares"
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

func DeleteProductControllers(c echo.Context) error {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "False Param",
		})
	}
	id_user_product, _ := databases.GetIDUserProduct(conv_id)
	log.Println("id_user_product", id_user_product)
	logged := middlewares.ExtractTokenId(c)
	log.Println("idlogged", logged)
	if uint(logged) != id_user_product {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Access Forbidden",
		})
	}
	_, e := databases.DeleteProduct(conv_id)

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
