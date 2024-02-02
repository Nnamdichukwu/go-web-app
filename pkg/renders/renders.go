package renders
import (
	"html/template"
	"net/http"
	"fmt"
)
func RenderTemplate(w http.ResponseWriter, tmpl string){
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl,"./templates/base.layout.gohtml")
	if err := parsedTemplate.Execute(w, nil);  err != nil{
		fmt.Println("error parsing template:", err)
		return
	}
}