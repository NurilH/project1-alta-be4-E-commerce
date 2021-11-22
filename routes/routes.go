package routes

import (
	"project_altabe4_1/constants"
	"project_altabe4_1/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	e := echo.New()
	// users
	e.POST("/users", controllers.CreateUserControllers)
	e.POST("/login", controllers.LoginUserControllers)

	// group JWT
	j := e.Group("/jwt")
	j.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	// users
	j.GET("/users/:id", controllers.GetUserControllers)
	j.PUT("/users/:id", controllers.UpdateUserControllers)
	j.DELETE("/users/:id", controllers.DeleteUserControllers)

	// products
	j.GET("/products/:id", controllers.GetProductByIdControllers)
	j.DELETE("/products/:id", controllers.DeleteProductControllers)
	j.POST("/products", controllers.CreateProductControllers)
	j.GET("/products", controllers.GetProductsControllers)
	j.PUT("/products/:id", controllers.UpdateProductControllers)

	//cart
	j.POST("/cart", controllers.CreateCartControllers)
	j.PUT("/cart/:id", controllers.UpdateCartControllers)
	j.GET("/cart", controllers.GetAllCartControllers)
	j.DELETE("/cart/:id", controllers.DeleteCartControllers)

	//credit card
	j.POST("/credit", controllers.CreateCreditControllers)
	j.DELETE("/credit/:id", controllers.DeleteCreditControllers)

	//order
	j.POST("/order", controllers.CreateOrderControllers)

	j.GET("/order", controllers.GetOrderControllers)
	return e
}
