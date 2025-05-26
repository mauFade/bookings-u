package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mauFade/bookings-u/internal/config"
	"github.com/mauFade/bookings-u/internal/forms"
	"github.com/mauFade/bookings-u/internal/models"
	"github.com/mauFade/bookings-u/internal/renders"
)

type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

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

	renders.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	strMap := map[string]string{
		"text": "Hello, this is the about page",
	}

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	strMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: strMap,
	})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
	})
}

func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	form.Has("first_name", r)
	form.Has("last_name", r)
	form.Has("email", r)
	form.Has("phone", r)

	if !form.Valid() {
		data := make(map[string]any)
		data["reservation"] = reservation

		renders.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
	}
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("Start date %s, end date %s", start, end)))
}

func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	res := jsonResponse{
		Ok:      true,
		Message: "available!",
	}

	output, err := json.MarshalIndent(res, "", "    ")

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}
