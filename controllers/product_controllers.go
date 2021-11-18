package controllers

import (
	"net/http"
	"project_altabe4_1/lib/databases"
	"project_altabe4_1/middlewares"
	"project_altabe4_1/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetProductsController(c echo.Context) error {
	products, err := databases.GetAllProduct()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Successful Operation",
		"data":    products,
	})
}

func CreateProductControllers(c echo.Context) error {
	Product := models.Product{}
	c.Bind(&Product)
	logged := middlewares.ExtractTokenId(c)
	Product.UsersID = uint(logged)
	_, e := databases.CreateProduct(&Product)
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

func UpdateProductControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "False Param",
		})
	}

	logged := middlewares.ExtractTokenId(c)
	if logged != id {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Access Forbidden",
		})
	}

	product := models.Product{}
	c.Bind(&product)
	_, err = databases.UpdateProduct(id, &product)
	if err != nil {
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
