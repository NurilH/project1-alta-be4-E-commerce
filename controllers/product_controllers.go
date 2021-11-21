package controllers

import (
	"net/http"
	"project_altabe4_1/lib/databases"
	"project_altabe4_1/middlewares"
	"project_altabe4_1/models"
	"project_altabe4_1/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

// controller untuk membuat produk
func CreateProductControllers(c echo.Context) error {
	Product := models.Product{}
	c.Bind(&Product)
	logged := middlewares.ExtractTokenId(c)
	Product.UsersID = uint(logged)
	_, e := databases.CreateProduct(&Product)
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

// controller untuk menampilkan seluruh produk
func GetProductsControllers(c echo.Context) error {
	products, err := databases.GetAllProduct()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(products))
}

// controller untuk menampilkan produk by id
func GetProductByIdControllers(c echo.Context) error {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	product, e := databases.GetProductById(conv_id)
	if e != nil || product == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(product))
}

// controller untuk menghapus produk by id
func DeleteProductControllers(c echo.Context) error {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	id_user_product, _ := databases.GetIDUserProduct(conv_id)
	logged := middlewares.ExtractTokenId(c)
	if uint(logged) != id_user_product {
		return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
	}
	databases.DeleteProduct(conv_id)
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

// controller untuk memperbarui data produk by id
func UpdateProductControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	id_user_product, _ := databases.GetIDUserProduct(id)
	logged := middlewares.ExtractTokenId(c)

	if logged != int(id_user_product) {
		return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
	}

	product := models.Product{}
	c.Bind(&product)
	databases.UpdateProduct(id, &product)
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}
