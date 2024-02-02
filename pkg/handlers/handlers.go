package handlers 
import (
	"net/http"
	"github.com/Nnamdichukwu/go-web-app/pkg/renders"
)
func Home(w http.ResponseWriter, r *http.Request){
	renders.RenderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request){
	renders.RenderTemplate(w, "about.page.html")
	
}


