package handlers

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		NotFound(w, r)
		return
	}
	fmt.Println("Path: ", r.URL.Path)
	fmt.Println("Method: ", r.Method)

	fmt.Fprintln(w, "Hello world!")
}
