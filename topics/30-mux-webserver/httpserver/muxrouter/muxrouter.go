package muxrouter

import (
	"net/http"

	"github.com/AntonioMartinezFernandez/udemy-golang-profesional/topics/30-mux-webserver/httpserver/handlers"
)

func CreateMuxRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/notfound", handlers.NotFound)
	mux.HandleFunc("/other", handlers.OtherHandler)
	mux.HandleFunc("/parameters", handlers.ParametersHandler)

	return mux
}
