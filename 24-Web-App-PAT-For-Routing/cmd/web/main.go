package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/VargheseVibin/GoWebPackage01/pkg/config"
	"github.com/VargheseVibin/GoWebPackage01/pkg/handlers"
	"github.com/VargheseVibin/GoWebPackage01/pkg/render"
)

const portNumber = ":8081"

func main() {
	var appConf config.AppConfig
	var err error
	appConf.TemplateCache, err = render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	appConf.UseCache = false
	repo := handlers.NewRepo(&appConf)
	handlers.NewHandlers(repo)
	render.NewTemplates(&appConf)

	fmt.Printf("This is a sample line")

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting Application on Port %v", portNumber)

	// http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&appConf),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
