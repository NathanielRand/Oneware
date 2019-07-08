package main

import (
	"net/http"

	"oneware/views"

	"github.com/gorilla/mux"
)

var aboutView *views.View
var contactView *views.View
var homeView *views.View
var partnersView *views.View
var projectsView *views.View
var servicesView *views.View

func about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := aboutView.Template.ExecuteTemplate(w,
		aboutView.Layout, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Template.ExecuteTemplate(w,
		contactView.Layout, nil); err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Template.ExecuteTemplate(w,
		homeView.Layout, nil); err != nil {
		panic(err)
	}
}

func partners(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := partnersView.Template.ExecuteTemplate(w,
		partnersView.Layout, nil); err != nil {
		panic(err)
	}
}

func projects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := projectsView.Template.ExecuteTemplate(w,
		projectsView.Layout, nil); err != nil {
		panic(err)
	}
}

func services(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := servicesView.Template.ExecuteTemplate(w,
		servicesView.Layout, nil); err != nil {
		panic(err)
	}
}

func main() {
	aboutView = views.NewView("materialize", "views/about.gohtml")
	contactView = views.NewView("materialize", "views/contact.gohtml")
	homeView = views.NewView("materialize", "views/home.gohtml")
	partnersView = views.NewView("materialize", "views/partners.gohtml")
	projectsView = views.NewView("materialize", "views/projects.gohtml")
	servicesView = views.NewView("materialize", "views/services.gohtml")

	r := mux.NewRouter()

	// Asset Routes
	assetHandler := http.FileServer(http.Dir("./assets/"))
	assetHandler = http.StripPrefix("/assets/", assetHandler)
	r.PathPrefix("/assets/").Handler(assetHandler)

	// Static Routes
	r.HandleFunc("/", home)
	r.HandleFunc("/about", about)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/partners", partners)
	r.HandleFunc("/projects", projects)
	r.HandleFunc("/services", services)
	http.ListenAndServe(":3000", r)
}
