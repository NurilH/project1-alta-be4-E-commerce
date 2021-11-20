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

func CreateCreditControllers(c echo.Context) error {
	Credit := models.Credit{}
	c.Bind(&Credit)
	logged := middlewares.ExtractTokenId(c)
	Credit.UsersID = uint(logged)

	_, e := databases.CreateCredit(&Credit)
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

func DeleteCreditControllers(c echo.Context) error {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "False Param",
		})
	}
	id_user_credit, _ := databases.GetIDUserCredit(conv_id)
	log.Println("id_user_credit", id_user_credit)
	logged := middlewares.ExtractTokenId(c)
	log.Println("idlogged", logged)
	if uint(logged) != id_user_credit {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Access Forbidden",
		})
	}
	databases.DeleteCredit(conv_id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Successful Operation",
	})
}
