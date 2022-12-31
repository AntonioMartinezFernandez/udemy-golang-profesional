package router

import (
	"net/http"

	"github.com/AntonioMartinezFernandez/udemy-golang-profesional/topics/29-webserver/httpserver/handlers"
)

func CreateRouter() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/notfound", handlers.NotFound)
	http.HandleFunc("/other", handlers.OtherHandler)
	http.HandleFunc("/parameters", handlers.ParametersHandler)
}
