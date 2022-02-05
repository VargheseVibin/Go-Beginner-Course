package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/VargheseVibin/GoWebPackage01/pkg/config"
	"github.com/VargheseVibin/GoWebPackage01/pkg/models"
)

var functions = template.FuncMap{}
var appConf *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	appConf = a
}

// Function to add any default data into the template
// Placeholder fucntion for now without any logic
func AddDefaultDataToTemplate(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var templCache map[string]*template.Template
	if appConf.UseCache {
		templCache = appConf.TemplateCache
	} else {
		templCache, _ = CreateTemplateCache()
	}

	templateForRender, ok := templCache[tmpl]
	if !ok {
		log.Fatal("No Template " + tmpl + " found in the TemplateCache")
	}
	// Here byteBuffer is created to read-in the template into the buffer.
	// Unlike, in the last project, where the template was directly read from the disk
	byteBuffer := new(bytes.Buffer)
	AddDefaultDataToTemplate(td)
	_ = templateForRender.Execute(byteBuffer, td)
	_, err := byteBuffer.WriteTo(w)
	if err != nil {
		fmt.Println("Error wiritng template to the browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		baseName := filepath.Base(page)
		fmt.Println("Render looking for page", baseName)
		ts, err := template.New(baseName).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		layoutMatches, err1 := filepath.Glob("./templates/*.layout.gohtml")
		if err1 != nil {
			return myCache, err1
		}
		if len(layoutMatches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return myCache, err
			}
		}
		myCache[baseName] = ts
	}
	return myCache, nil

}
