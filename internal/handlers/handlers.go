package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/AdirNoyman/bookings/internal/config"
	"github.com/AdirNoyman/bookings/internal/models"
	"github.com/AdirNoyman/bookings/internal/render"
	"log"
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

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Adiros 🤓"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		// The data I'm might pass in the template
		StringMap: stringMap,
	})
}

// Reservation renders the make a reservation page and displays the reservation form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})

}

// President renders the president page
func (m *Repository) President(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "president.page.tmpl", &models.TemplateData{})

}

// RoyalSweet renders the royal-sweet page
func (m *Repository) RoyalSweet(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "royal-sweet.page.tmpl", &models.TemplateData{})

}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})

}

// PostAvailability renders the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {

	// Get the form's data that was submitted by the user
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	_, err := w.Write([]byte(fmt.Sprintf("Start date is %s and the End date is %s 🤓", start, end)))
	if err != nil {
		return
	}

}

// AvailabilityJSON handles requests for availability and sends JSON response
type jsonResponse struct {
	// In order to parse this Go object as JSON, we need the name of the variable to start in Upper case and to set what it's name should be in the json file
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {

	resp := jsonResponse{

		OK:      false,
		Message: "Available",
	}

	// Marshal = turn Golang object into json
	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println(err)
	}

	// 1. Write bytes that will form the json that will be passed to the handlers response
	// 2. Setting the header to say what type of response I'm sending
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// MakeReservation renders the make-reservation page
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})

}

// Contact renders the search availability page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})

}
