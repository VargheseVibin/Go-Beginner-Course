package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/VargheseVibin/GoWebPackage01/pkg/config"
	"github.com/VargheseVibin/GoWebPackage01/pkg/handlers"
	"github.com/VargheseVibin/GoWebPackage01/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8081"

var appConf config.AppConfig
var session *scs.SessionManager

func main() {
	appConf.InProduction = false // set this as true in Production

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConf.InProduction

	appConf.Session = session

	var err error
	appConf.TemplateCache, err = render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	appConf.UseCache = true
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
