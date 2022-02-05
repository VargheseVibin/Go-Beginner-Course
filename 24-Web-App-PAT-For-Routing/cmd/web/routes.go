package main

import (
	"net/http"

	"github.com/VargheseVibin/GoWebPackage01/pkg/config"
	"github.com/VargheseVibin/GoWebPackage01/pkg/handlers"
	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	// http handler called Multiplexer(mux)
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux
}
