package handlers

import (
	"net/http"

	"github.com/VargheseVibin/GoWebPackage01/pkg/render"
)

// Home if the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
