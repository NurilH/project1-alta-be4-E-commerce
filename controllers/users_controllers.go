package controllers

import (
	"net/http"
	"project_altabe4_1/lib/databases"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUsersControllers(c echo.Context) error {

}

func CreateUserController(c echo.Context) error {

}

func DeleteUserController(c echo.Context) error {

}

func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "Bad Request",
		})

	}
	user, e := databases.GetUser(id)
	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "Bad Request",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "succes",
		"user":    user,
	})

}
