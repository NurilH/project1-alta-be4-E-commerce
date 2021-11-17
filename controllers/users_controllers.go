package controllers

import (
	"net/http"
	"project_altabe4_1/lib/databases"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUserControllers(c echo.Context) error {
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request",
		})
	}
	user, e := databases.GetUser(convId)
	if e != nil || user == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Bad Request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Successful Operation",
		"data":    user,
	})
}

func CreateUserController(c echo.Context) error {

}

func DeleteUserController(c echo.Context) error {

}

func UpdateUserController(c echo.Context) error {

}
