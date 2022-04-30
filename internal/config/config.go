package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

// AppConfig holds the application configuration
type AppConfig struct {
	// UseCache - If we are in production will use the template cache. But if we are in development, we will re-build it every time we load the app
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	// InProduction am I in production environment (used by the session object)
	InProduction bool
	// Session declaring our session as a global variable so our app could use it in different places
	Session *scs.SessionManager
}
