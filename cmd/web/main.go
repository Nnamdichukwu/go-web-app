package main

import (
	"fmt"
	"net/http"
	"github.com/Nnamdichukwu/go-web-app/pkg/handlers"
	
)
const portNumber = ":8080"

func main()  {
	
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about",handlers.About)

	fmt.Println("Starting application on port ", portNumber)

	http.ListenAndServe(portNumber,nil)
}