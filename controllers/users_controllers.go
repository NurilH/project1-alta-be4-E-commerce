package controllers

import (
	"net/http"
	"project_altabe4_1/lib/databases"
	"project_altabe4_1/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUserControllers(c echo.Context) error {
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "False Param",
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

func CreateUserControllers(c echo.Context) error {
	new_user := models.Users{}
	c.Bind(&new_user)
	_, e := databases.CreateUser(&new_user)
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

// func DeleteUserControllers(c echo.Context) error {

// }

// func UpdateUserControllers(c echo.Context) error {

// }
