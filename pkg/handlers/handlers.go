package handlers

import (
	"net/http"

	"github.com/Nnamdichukwu/go-web-app/pkg/config"
	"github.com/Nnamdichukwu/go-web-app/pkg/models"
	"github.com/Nnamdichukwu/go-web-app/pkg/renders"
)

//Repo is the repository used by the handlers
var Repo *Repository

//Repository is the Repository type
type Repository struct{
	App *config.AppConfig
}
//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App: a,
	}
}
//NewHandlers sets the repository for the handlers 
func NewHandlers (r *Repository){
	Repo = r
}
func (m *Repository) Home(w http.ResponseWriter, r *http.Request){
	renders.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request){
	
	stringMap := map[string]string {}
	stringMap["test"] = "Hello Again"
	renders.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}


