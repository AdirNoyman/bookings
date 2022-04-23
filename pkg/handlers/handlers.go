package handlers

import (
	"github.com/AdirNoyman/bookings/pkg/config"
	"github.com/AdirNoyman/bookings/pkg/models"
	"github.com/AdirNoyman/bookings/pkg/render"
	"net/http"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repo is the repository used by the handlers
var Repo *Repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{

		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {

	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	// catch the users IP address and store it in the session object
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Adiros 🤓"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		// The data I'm might pass in the template
		StringMap: stringMap,
	})
}

// Reservation renders the make a reservation page and displays the reservation form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})

}

// President renders the president page
func (m *Repository) President(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "president.page.tmpl", &models.TemplateData{})

}

// RoyalSweet renders the royal-sweet page
func (m *Repository) RoyalSweet(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "royal-sweet.page.tmpl", &models.TemplateData{})

}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})

}

// Contact renders the search availability page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})

}
