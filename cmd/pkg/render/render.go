package render

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/dineshpawar22/booking/cmd/pkg/config"
	"github.com/dineshpawar22/booking/cmd/pkg/models"
)

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplates(w http.ResponseWriter, page string, td *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[page]

	fmt.Println("pages:", t)
	if !ok {
		fmt.Println("Unable to get template from cache :", ok)
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing templates to browser : ", err)
	}
}

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates
func NewTemplates(a *config.AppConfig) {
	app = a
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		//Test
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {

			ts, err = ts.ParseGlob("./templates/*.layout.html")

			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}

	return myCache, nil

}
