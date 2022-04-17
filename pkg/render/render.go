package render

import (
	"bytes"
	"fmt"
	"github.com/AdirNoyman/bookings/pkg/config"
	"github.com/AdirNoyman/bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Preparation for custom functions I might insert to the templates
var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {

	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	// td = data I want that will be available on every page
	return td
}

// RenderTemplate Render the html templates
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template

	// If we are in production will use the template cache. But if we are in development, we will re-build it every time we load the app
	if app.UseCache {
		// Get the template cache variable from the app config (that was created in main.go)
		tc = app.TemplateCache

	} else {

		tc, _ = CreateTemplateCache()

	}

	t, ok := tc[tmpl]

	// Checking if the name of the template that was passed is in the template cache
	if !ok {

		log.Fatal("There is an error.Could not get template from the template cache. Killing the app 😫")
	}

	// Turn the parsed template we found in to bytes
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)
	// Storing the bytes we got in to this buf variable
	_ = t.Execute(buf, td)

	// Writing the template into the browser
	_, err := buf.WriteTo(w)

	if err != nil {

		fmt.Println("Error writing template to browser", err)
	}

	//parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	//err = parsedTemplate.Execute(w, nil)
	//
	//if err != nil {
	//
	//	fmt.Println("Error parsing template 🥴:", err)
	//	return
	//}

}

// CreateTemplateCache creates templates cache as a map(name of template : parsed template)
func CreateTemplateCache() (map[string]*template.Template, error) {

	// This creates a map where the string is the name(key) of the template and the value will the parsed template
	myCache := map[string]*template.Template{}

	// Get all the files in the templates folder that start with the word page
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {

		return myCache, err
	}

	// _ = index, page = template name
	for _, page := range pages {

		name := filepath.Base(page)

		// ts = template set
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {

			return myCache, err
		}

		// Look in the templates for any file that contains layout
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {

			return myCache, err
		}

		if len(matches) > 0 {

			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
		}

		myCache[name] = ts
	}

	return myCache, nil

}
