package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

var films = map[string][]Film{
	"Films": {
		{Title: "The Godfather", Director: "Francis Ford Coppola"},
		{Title: "Blade Runner", Director: "Ridley Scott"},
		{Title: "The Thing", Director: "John Carpenter"},
	},
}

type Film struct {
	Title    string
	Director string
}

func main() {
	log.Println("Server running on port 8000...")

	http.HandleFunc("/", handleFunc)
	http.HandleFunc("/add-film/", handleAddFilm)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handleFunc(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, films)
}

func handleAddFilm(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
}
