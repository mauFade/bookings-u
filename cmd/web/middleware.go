package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func NoSurf(next http.Handler) http.Handler {
	ns := nosurf.New(next)

	ns.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return ns
}
