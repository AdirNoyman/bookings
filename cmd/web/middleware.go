package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

// NoSurf creates no surf token that protects all POST requests
func NoSurf(next http.Handler) http.Handler {

	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{

		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad - turns our server from stateless (default mode) to stateful by saving the session data that is in the cookie
// LoadAndSave provides middleware which automatically loads and saves session
// data for the current request, and communicates the session token to and from
// the client in a cookie.
func SessionLoad(next http.Handler) http.Handler {

	return session.LoadAndSave(next)
}
