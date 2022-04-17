package main

import (
	"fmt"
	"github.com/AdirNoyman/bookings/pkg/config"
	"github.com/AdirNoyman/bookings/pkg/handlers"
	"github.com/AdirNoyman/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

// Declare our configuration for this app
var app config.AppConfig

var session *scs.SessionManager

func main() {

	// Change this to rue when in production
	app.InProduction = false

	// Create a session object
	session = scs.New()
	// Configure how long will the session data be held (e.g. 24h)
	session.Lifetime = 24 * time.Hour
	// By default the data is saved in a session cookie
	// Configure the session to persist if the user closes the browser window
	session.Cookie.Persist = true
	// Configuring the cookie to apply only for this site
	session.Cookie.SameSite = http.SameSiteLaxMode
	// Configuring the cookie not to be encrypted and to use http (dev only!)
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can't create template cache 😩")
	}

	app.TemplateCache = tc

	// Because I'm in development mode, I'm re-building the templates every time the page reloads
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	// Server
	fmt.Println(fmt.Sprintf("Starting application on port %s 😎🤟", portNumber))
	//_ = http.ListenAndServe(portNumber, nil)

	// srv = server
	// Mount our routes to our server
	srv := &http.Server{

		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()

	// If there is an error -> Log it
	log.Fatal(err)
}
