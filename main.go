package main

import (
	"log"
	"net/http"
	"project_altabe4_1/config"
	"project_altabe4_1/middlewares"
	"project_altabe4_1/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	middlewares.LogMiddlewares(e)
	if err := e.StartTLS(":443", "server.crt", "server.key"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
