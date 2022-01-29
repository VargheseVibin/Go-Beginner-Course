package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	templCache, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	templateForRender, ok := templCache[tmpl]
	if !ok {
		log.Fatal(err)
	}
	// Here byteBuffer is created to read-in the template into the buffer.
	// Unlike, in the last project, where the template was directly read from the disk
	byteBuffer := new(bytes.Buffer)
	_ = templateForRender.Execute(byteBuffer, nil)
	_, err = byteBuffer.WriteTo(w)
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
