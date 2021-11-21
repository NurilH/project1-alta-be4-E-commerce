package controllers

import (
	"log"
	"net/http"
	"project_altabe4_1/helper"
	"project_altabe4_1/lib/databases"
	"project_altabe4_1/middlewares"
	"project_altabe4_1/models"
	"project_altabe4_1/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUserControllers(c echo.Context) error {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)
	log.Println("id", conv_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	user, e := databases.GetUser(conv_id)
	if e != nil || user == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(user))
}

func CreateUserControllers(c echo.Context) error {
	new_user := models.Users{}
	c.Bind(&new_user)
	new_user.Password, _ = helper.HashPassword(new_user.Password)
	_, e := databases.CreateUser(&new_user)
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(new_user))
}

func DeleteUserControllers(c echo.Context) error {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}

	logged := middlewares.ExtractTokenId(c)
	if logged != conv_id {
		return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
	}

	_, e := databases.DeleteUser(conv_id)
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

func UpdateUserControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}

	logged := middlewares.ExtractTokenId(c)
	if logged != id {
		return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
	}
	users := models.Users{}
	c.Bind(&users)
	users.Password, _ = helper.HashPassword(users.Password)
	_, e := databases.UpdateUser(id, &users)
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

func LoginUserControllers(c echo.Context) error {
	user := models.Users{}
	c.Bind(&user)
	plan_pass := user.Password
	log.Println(plan_pass)
	token, e := databases.LoginUser(plan_pass, &user)
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.LoginFailedResponse())
	}
	return c.JSON(http.StatusOK, response.LoginSuccessResponse(token))
}

func GetUserControllersTesting() echo.HandlerFunc {
	return GetUserControllers
}

func UpdateUserControllersTesting() echo.HandlerFunc {
	return UpdateUserControllers
}

func DeleteUserControllersTesting() echo.HandlerFunc {
	return DeleteUserControllers
}
