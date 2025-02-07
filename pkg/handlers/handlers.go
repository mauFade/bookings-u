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
