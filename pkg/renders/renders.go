package renders

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)
func RenderTemplateTest(w http.ResponseWriter, tmpl string){
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl,"./templates/base.layout.gohtml")
	if err := parsedTemplate.Execute(w, nil);  err != nil{
		fmt.Println("error parsing template:", err)
		return
	}
}

var tc = make(map[string]*template.Template)
func RenderTemplate(w http.ResponseWriter, t string){
	var tmpl *template.Template
	var err error
	_, inMap := tc[t]
	if !inMap{
		log.Println("creatong a new cache")
		err = createTemplateCache(t)
		if err != nil{
			log.Println(err)
		}
	}else {
		log.Println("using cached template")
	}
	tmpl = tc[t]
	err =  tmpl.Execute(w, nil)
	if err != nil{
		log.Println(err)
	}
	
}

func createTemplateCache(t string) error {
	templates := []string {
		fmt.Sprintf("./templates/%s",t),
		"./templates/base.layout.gohtml",
	}
	tmpl, err := template.ParseFiles(templates...)
	if err != nil{
		return err
	}
	tc[t] = tmpl
	return nil
}
