package handlers

import (
	"net/http"

	"github.com/mauFade/bookings-u/pkg/config"
	"github.com/mauFade/bookings-u/pkg/models"
	"github.com/mauFade/bookings-u/pkg/renders"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	renders.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	strMap := map[string]string{
		"text": "Hello, this is the about page",
	}

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	strMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: strMap,
	})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "generals.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "majors.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}
