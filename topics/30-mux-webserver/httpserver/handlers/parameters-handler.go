package handlers

import (
	"fmt"
	"net/http"
)

func ParametersHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/parameters" {
		NotFound(w, r)
		return
	}
	fmt.Println("Path: ", r.URL.Path)
	fmt.Println("Method: ", r.Method)

	/* Obtain parameters (method 1) */
	// name := r.FormValue("name")
	// surname := r.FormValue("surname")

	/* Obtain parameters (method 2) */
	parameters := r.URL.Query()

	fmt.Fprintf(w, "Hello, %s %s!", parameters.Get("name"), parameters.Get("surname"))
}
