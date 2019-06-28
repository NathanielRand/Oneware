package main

import (
	"net/http"

	"oneware/views"

	"github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View
var servicesView *views.View
var projectsView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Template.ExecuteTemplate(w,
		homeView.Layout, nil); err != nil {
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

func services(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := servicesView.Template.ExecuteTemplate(w,
		servicesView.Layout, nil); err != nil {
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

func main() {
	homeView = views.NewView("materialize", "views/home.gohtml")
	contactView = views.NewView("materialize", "views/contact.gohtml")
	servicesView = views.NewView("materialize", "views/services.gohtml")
	projectsView = views.NewView("materialize", "views/projects.gohtml")

	r := mux.NewRouter()

	// Asset Routes
	assetHandler := http.FileServer(http.Dir("./assets/"))
	assetHandler = http.StripPrefix("/assets/", assetHandler)
	r.PathPrefix("/assets/").Handler(assetHandler)

	// Static Routes
	r.HandleFunc("/", home)
	r.HandleFunc("/services", services)
	r.HandleFunc("/projects", projects)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}
