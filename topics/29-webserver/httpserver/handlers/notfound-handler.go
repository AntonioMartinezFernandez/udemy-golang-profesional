package handlers

import (
	"fmt"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Path: ", r.URL.Path)
	fmt.Println("Method: ", r.Method)

	http.Error(w, "Not found", http.StatusNotFound)
}
