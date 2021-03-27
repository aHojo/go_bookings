package handlers

import (
	"github.com/ahojo/go_bookings/pkg/Models"
	"github.com/ahojo/go_bookings/pkg/config"
	"github.com/ahojo/go_bookings/pkg/render"
	"net/http"
)

var Repo *Repository

// Repository is the respository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

/// NewHandlers sets the repository for the handers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the homepage handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.gohtml", &Models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Perform some logic

	stringMap := make(map[string]string)
	stringMap["test"] = "Hi Kairi, daddy loves you! "

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send data to the template
	render.RenderTemplate(w, "about.page.gohtml", &Models.TemplateData{
		StringMap: stringMap,
	})
}
