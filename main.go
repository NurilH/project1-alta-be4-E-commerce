package main

import (
	"project_altabe4_1/config"
	"project_altabe4_1/routes"
)

func main() {
	config.InitDB()
	e := routes.New()

	e.Logger.Fatal(e.Start(":8080"))
}
