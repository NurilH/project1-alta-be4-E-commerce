package routes

import (
	"project_altabe4_1/constants"
	"project_altabe4_1/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	e := echo.New()
	// e.GET("/users/:id", controllers.GetUserControllers)
	e.POST("/users", controllers.CreateUserControllers)
	// e.PUT("/users/:id", controllers.UpdateUserControllers)
	// e.DELETE("/users/:id", controllers.DeleteUserControllers)
	e.POST("/login", controllers.LoginUserControllers)

	// e.GET("/products/:id", controllers.GetProductByIdControllers)
	// e.DELETE("/products/:id", controllers.DeleteProductControllers)

	// group JWT
	j := e.Group("/jwt")
	j.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	j.GET("/users/:id", controllers.GetUserControllers)
	j.PUT("/users/:id", controllers.UpdateUserControllers)
	j.DELETE("/users/:id", controllers.DeleteUserControllers)

	//product
	j.GET("/products/:id", controllers.GetProductByIdControllers)
	j.DELETE("/products/:id", controllers.DeleteProductControllers)
	j.POST("/products", controllers.CreateProductControllers)
	j.GET("/products", controllers.GetProductsControllers)
	j.PUT("/products/:id", controllers.UpdateProductControllers)

	//cart
	j.POST("/cart", controllers.CreateCartControllers)

	return e
}
