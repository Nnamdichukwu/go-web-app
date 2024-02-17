package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Nnamdichukwu/go-web-app/internal/config"
	"github.com/Nnamdichukwu/go-web-app/internal/models"
	"github.com/Nnamdichukwu/go-web-app/internal/renders"
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
	remoteIP := r.RemoteAddr
	
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	renders.RenderTemplate(w, r, "home.page.gohtml", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request){
	
	stringMap := map[string]string {}
	stringMap["test"] = "Hello Again"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w,r, "about.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
	
}

func (m *Repository)Generals(w http.ResponseWriter, r *http.Request){
	
	renders.RenderTemplate(w, r, "generals.page.gohtml", &models.TemplateData{})
}

func (m *Repository)Majors(w http.ResponseWriter, r *http.Request){
	
	renders.RenderTemplate(w,r, "majors.page.gohtml", &models.TemplateData{})
}
func (m *Repository)Reservation(w http.ResponseWriter, r *http.Request){
	
	renders.RenderTemplate(w,r, "make-reservation.page.gohtml", &models.TemplateData{})
}
func (m *Repository)SearchAvailability(w http.ResponseWriter, r *http.Request){
	
	renders.RenderTemplate(w, r, "search-availability.page.gohtml", &models.TemplateData{})
}

func (m *Repository)PostAvailability(w http.ResponseWriter, r *http.Request){
	start := r.Form.Get("start")
	end := r.Form.Get("end") 
	 
	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s ", start, end)))

	
}
type jsonResponse struct{
	OK 		bool `json: "ok`
	Message string `json:"message"`
}
func (m *Repository)AvailabilityJSON(w http.ResponseWriter, r *http.Request){
	resp := jsonResponse{
		OK: true, 
		Message: "Available",
	}
	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println(err)
	}
	
	//This tells the browser what type of content i am sending
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

	
}
func (m *Repository)Contact(w http.ResponseWriter, r *http.Request){
	
	renders.RenderTemplate(w, r, "contact.page.gohtml", &models.TemplateData{})
}