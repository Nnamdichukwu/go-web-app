package renders

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Nnamdichukwu/go-web-app/internal/config"
	"github.com/Nnamdichukwu/go-web-app/internal/models"
	"github.com/justinas/nosurf"
)
var app *config.AppConfig
//NewTemplates sets the template of the app config
func NewTemplates(a *config.AppConfig){
	app = a
}

func AddDefaultData (td *models.TemplateData, r *http.Request) *models.TemplateData{
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td 
}
func RenderTemplate(w http.ResponseWriter,r *http.Request, tmpl string, td *models.TemplateData){
	//get the template cache from the app config
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	}else {
		tc, _= CreateTemplateCache()
		

	}
	t,ok := tc[tmpl]

	if !ok{
		log.Fatal("could not get template from template cache")
	}
	buf := new(bytes.Buffer)
	td = AddDefaultData(td, r)
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}
	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//get all of the files name *.page.gohtml from ./templates
	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil{
		return myCache, err
	}
	//range through all files ending with *.page.gohtml
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil{
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil{
			return myCache, err
		}
		if len(matches) > 0{
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return myCache, err
			}

		}
		myCache[name] = ts
	}
	return myCache, nil
}
