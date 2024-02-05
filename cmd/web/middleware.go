package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Hit the page", r.URL) 
		next.ServeHTTP(w,r)
	})
}

func NoSurf(next http.Handler)http.Handler{
	csrfHandler := nosurf.New(next)
	//this creates a cookie across the site
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}