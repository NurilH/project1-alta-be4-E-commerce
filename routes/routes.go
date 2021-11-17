package routes

import (
	"project_altabe4_1/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	e := echo.New()
	e.GET("/users/:id", controllers.GetUserControllers)
	e.POST("/users", controllers.CreateUserControllers)
	e.PUT("/users/:id", controllers.UpdateUserControllers)
	e.DELETE("/users/:id", controllers.DeleteUserControllers)
	return e
}
