package main

import (
	"log"
	"net/http"
	"text/template"
	"time"
)

func getWebsite(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	films := map[string][]Film{
		"Films": {
			{Title: "Finding Nemo", Director: "Me"},
			{Title: "Finding Dory", Director: "Not me"},
			{Title: "Finding me", Director: "Also not me"},
		},
	}
	tmpl.Execute(w, films)
}

func addFilmItem(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(w, "film-list-element", Film{
		Title:    title,
		Director: director,
	})
}

func main() {
	http.HandleFunc("/", getWebsite)
	http.HandleFunc("/add-film/", addFilmItem)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
