package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"project_altabe4_1/config"
	"project_altabe4_1/middlewares"
	"project_altabe4_1/routes"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	config.InitDB()
	e := routes.New()
	middlewares.LogMiddlewares(e)
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("altaproject.cloudns.ph"), //Your domain here
		Cache:      autocert.DirCache("certs"),                       //Folder for storing certificates
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	server := &http.Server{
		Addr: ":https",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	go http.ListenAndServe(":http", certManager.HTTPHandler(nil))

	log.Fatal(server.ListenAndServeTLS("", ""))
}
