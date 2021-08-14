package handlers

import (
	"net/http"

	"github.com/dineshpawar22/booking/cmd/pkg/config"
	"github.com/dineshpawar22/booking/cmd/pkg/models"
	"github.com/dineshpawar22/booking/cmd/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplates(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) Major_suit(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "major.page.html", &models.TemplateData{})
}

func (m *Repository) Generals_quarter(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "generals.page.html", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "search-availability.page.html", &models.TemplateData{})
}
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "contact.page.html", &models.TemplateData{})
}
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "make-reservation.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Dinesh 2"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplates(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
