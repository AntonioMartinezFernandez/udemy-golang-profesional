package httpserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AntonioMartinezFernandez/udemy-golang-profesional/topics/31-templates/httpserver/muxrouter"
)

func Start() {
	// Mux Router
	mux := muxrouter.CreateMuxRouter()

	// Static Files
	staticFiles := http.FileServer(http.Dir("./topics/31-templates/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFiles))

	// Http server
	server := &http.Server{
		Addr:    "0.0.0.0:3000",
		Handler: mux,
	}

	fmt.Println("Server running on port 3000")
	log.Fatal(server.ListenAndServe())

}
