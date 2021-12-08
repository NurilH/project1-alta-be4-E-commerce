package main

import (
	"net/http"
	"project_altabe4_1/config"
	"project_altabe4_1/middlewares"
	"project_altabe4_1/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	config.InitDB()
	e := routes.New()
	middlewares.LogMiddlewares(e)
	// e.Logger.Fatal(e.Start(":8080"))
	// if err := e.StartTLS(":443", "server.crt", "server.key"); err != http.ErrServerClosed {
	// 	log.Fatal(err)
	// }
	// e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("<DOMAIN>")
	// Cache certificates to avoid issues with rate limits (https://letsencrypt.org/docs/rate-limits)

	e.AutoTLSManager.Cache = autocert.DirCache("/home/ubuntu/.cache")
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
			<h1>Welcome to Echo!</h1>
			<h3>TLS certificates automatically installed from Let's Encrypt :)</h3>
		`)
	})

	e.Logger.Fatal(e.StartAutoTLS(":443"))
}
