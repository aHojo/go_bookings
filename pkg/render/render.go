package render

import (
	"bytes"
	"fmt"
	"github.com/ahojo/go_bookings/pkg/Models"
	"github.com/ahojo/go_bookings/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package.
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *Models.TemplateData) *Models.TemplateData {
	return td
}
func RenderTemplate(w http.ResponseWriter, tmpl string, data *Models.TemplateData) {
	var tc map[string]*template.Template
	// if UseCache is true, we are in prod, and don't want to read from disk every time
	// otherwise, we want to be able to see changes to our html.
	if app.UseCache {
		// get the template cache, from the app config.
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatalln("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	data = AddDefaultData(data)
	_ = t.Execute(buf, data)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser")
	}
}

// CreateTemplateCache creates a template cache as a map.
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return myCache, err
			}

		}

		myCache[name] = ts
	}
	return myCache, nil
}
