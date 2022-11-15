package handlers

import (
	"github.com/andres15mol/booking/pkg/config"
	"github.com/andres15mol/booking/pkg/models"
	"github.com/andres15mol/booking/pkg/render"

	"net/http"
)

//Repo the repository used by the handlers
var Repo *Repository

//Repository is the repository
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App:a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository){
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// General is the handler for the General page
func (m *Repository) General(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "general.page.tmpl", &models.TemplateData{})
}

// Major is the handler for the Major's Suite page
func (m *Repository) Major(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "major.page.tmpl", &models.TemplateData{})
}
// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//performe some logic

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp
	//send some data to the template
	
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Contact is the handler for the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}

// Reservation is the handler for the book page
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}

// Reservation is the handler for the book page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "reservation.page.tmpl", &models.TemplateData{})
}