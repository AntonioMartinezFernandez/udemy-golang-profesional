package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		NotFound(w, r)
		return
	}
	fmt.Println("Path: ", r.URL.Path)
	fmt.Println("Method: ", r.Method)

	// Create user data
	userData := Users{
		Name:    "Antonio Stradivari",
		Age:     379,
		Active:  true,
		Admin:   true,
		Courses: []Course{{"Golang"}, {"NodeJS"}},
	}

	// Map of functions
	functionsMap := template.FuncMap{
		"sayHello": sayHello,
	}

	// Create template (basic)
	// t, err := template.ParseFiles("./topics/31-templates/html/index.tmpl")

	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// } else {
	// 	t.Execute(w, userData)
	// }

	// Create template (complete)
	// t := template.Must(template.New("index.tmpl").Funcs(functionsMap).ParseFiles("./topics/31-templates/html/index.tmpl", "./topics/31-templates/html/base.tmpl"))

	// Create template (glob mode)
	t := template.Must(template.New("index.tmpl").Funcs(functionsMap).ParseGlob("./topics/31-templates/html/*.tmpl"))

	t.Execute(w, userData)
}

// Types
type Users struct {
	Name    string
	Age     int
	Active  bool
	Admin   bool
	Courses []Course
}

type Course struct {
	Title string
}

// Functions
func sayHello() string {
	return "Hello from function!"
}
