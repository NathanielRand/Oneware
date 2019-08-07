package main

import (
	"net/http"

	"Oneware/views"

	"github.com/gorilla/mux"
)

var notFoundView *views.View
var aboutView *views.View
var contactView *views.View
var homeView *views.View
var partnersView *views.View
var projectsView *views.View
var tutorialsView *views.View
var servicesView *views.View
var stackView *views.View

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := notFoundView.Template.ExecuteTemplate(w,
		notFoundView.Layout, nil); err != nil {
		panic(err)
	}
}

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

func tutorials(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tutorialsView.Template.ExecuteTemplate(w,
		tutorialsView.Layout, nil); err != nil {
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

func stack(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := stackView.Template.ExecuteTemplate(w,
		stackView.Layout, nil); err != nil {
		panic(err)
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/favicon.ico")
}

func main() {
	notFoundView = views.NewView("materialize", "views/notfound.gohtml")
	aboutView = views.NewView("materialize", "views/about.gohtml")
	contactView = views.NewView("materialize", "views/contact.gohtml")
	homeView = views.NewView("materialize", "views/home.gohtml")
	partnersView = views.NewView("materialize", "views/partners.gohtml")
	projectsView = views.NewView("materialize", "views/projects.gohtml")
	tutorialsView = views.NewView("materialize", "views/tutorials.gohtml")
	servicesView = views.NewView("materialize", "views/services.gohtml")
	stackView = views.NewView("materialize", "views/stack.gohtml")

	http.HandleFunc("/favicon.ico", faviconHandler)

	r := mux.NewRouter()

	// 404 Not Found
	r.NotFoundHandler = http.HandlerFunc(notFound)

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
	r.HandleFunc("/tutorials", tutorials)
	r.HandleFunc("/services", services)
	r.HandleFunc("/stack", stack)
	http.ListenAndServe(":8080", r)
}
