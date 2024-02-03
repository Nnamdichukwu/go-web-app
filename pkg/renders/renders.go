package renders

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)


func RenderTemplate(w http.ResponseWriter, tmpl string){
	//get the template cache from the app config
	tc, err := CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}
	t,ok := tc[tmpl]

	if !ok{
		log.Fatal()
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
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
