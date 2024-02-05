package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Nnamdichukwu/go-web-app/pkg/config"
	"github.com/Nnamdichukwu/go-web-app/pkg/handlers"
	"github.com/Nnamdichukwu/go-web-app/pkg/renders"
	"github.com/alexedwards/scs/v2"
)
const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager
func main()  {
	
	// change this to true when in production
	app.InProduction = false
	session =  scs.New()
	session.Lifetime = 24 * time.Hour // this will create a session that lasts for 24 hours
	session.Cookie.Persist = true 
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction 

	app.Session = session
	tc, err := renders.CreateTemplateCache()
	if err != nil{
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc 
	app.UseCache = false 
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	renders.NewTemplates(&app)
	

	fmt.Println("Starting application on port ", portNumber)

	serve := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}
	err = serve.ListenAndServe()
	log.Fatal(err)
}