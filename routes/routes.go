package routes

import (
	"project_altabe4_1/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	e := echo.New()
	e.GET("/users/:id", controllers.GetUserControllers)
	e.PUT("/users/:id", controllers.UpdateUserController)
	return e
}
