package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	template := template.Must(template.ParseFiles("templates/base.html"))
	template.Execute(w, map[string]int{"currentPage": 1})
}

func page_2Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/page_2" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	template := template.Must(template.ParseFiles("templates/base.html"))
	template.Execute(w, map[string]int{"currentPage": 2})
}

func page_3Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/page_3" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	template := template.Must(template.ParseFiles("templates/base.html"))
	template.Execute(w, map[string]int{"currentPage": 3})
}
func page_4Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/page_4" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	template := template.Must(template.ParseFiles("templates/base.html"))
	template.Execute(w, map[string]int{"currentPage": 4})
}
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		tmpl, _ := template.ParseFiles("templates/404.html")
		_ = tmpl.Execute(w, nil)
	}
}

func main() {
	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))

	// Handles routing:
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/page_2", page_2Handler)
	http.HandleFunc("/page_3", page_3Handler)
	http.HandleFunc("/page_4", page_4Handler)

	// Launches the server:
	preferredPort := ":8080"
	fmt.Printf("Starting server at port %v\n", preferredPort)
	if err := http.ListenAndServe(preferredPort, nil); err != nil {
		log.Fatal(err)
	}
}
