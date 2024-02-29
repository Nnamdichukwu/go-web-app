package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Nnamdichukwu/go-web-app/internal/config"
	"github.com/Nnamdichukwu/go-web-app/internal/forms"
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
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation
	renders.RenderTemplate(w,r, "make-reservation.page.gohtml", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}
func (m *Repository)PostReservation(w http.ResponseWriter, r *http.Request){
	 if err := r.ParseForm(); err != nil{
		log.Println(err)
		return
	 }

	 reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName: r.Form.Get("last_name"),
		Email: r.Form.Get("email"),
		Phone: r.Form.Get("phone"),
	 }

	 form := forms.New(r.PostForm)

	 //form.Has("first_name",r)
	 form.Required("first_name", "last_name", "email", "phone")
	 form.MinLength("first_name", 3, r)
	 form.IsEmail("email")
	 if !form.Valid(){
		data := make(map[string]interface{})
		data["reservation"] = reservation
		renders.RenderTemplate(w,r, "make-reservation.page.gohtml", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	 }
	 m.App.Session.Put(r.Context(),"reservation", reservation)

	http.Redirect(w,r,"/reservation-summary",http.StatusSeeOther)
}
func (m *Repository)ReservationSummary(w http.ResponseWriter, r *http.Request){
	reservation, ok  := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok{
		log.Println("Cannot get item from sessions")
		m.App.Session.Put(r.Context(),"error","Can't get reservation from session")
		http.Redirect(w,r,"/", http.StatusTemporaryRedirect)
		return
	}
	
	m.App.Session.Remove(r.Context(),"reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation
	renders.RenderTemplate(w,r,"reservation-summary.page.gohtml",&models.TemplateData{
		Data: data, 
	})
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