package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Nnamdichukwu/go-web-app/pkg/config"
	"github.com/Nnamdichukwu/go-web-app/pkg/handlers"
	"github.com/Nnamdichukwu/go-web-app/pkg/renders"
)
const portNumber = ":8080"

func main()  {
	var app config.AppConfig

	tc, err := renders.CreateTemplateCache()
	if err != nil{
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc 
	app.UseCache = false 
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	renders.NewTemplates(&app)
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about",handlers.Repo.About)

	fmt.Println("Starting application on port ", portNumber)

	http.ListenAndServe(portNumber,nil)
}