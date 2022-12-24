package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

func writeToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the next page")
		next.ServeHTTP(w, r)
	})
}

// add csrftoekn to every request
func noSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// session loads and saves the question on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
