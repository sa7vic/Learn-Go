package main

import (
	"html/template"
	"net/http"
	"strconv"
	"sync"
)

type Note struct {
	ID      int
	Title   string
	Content string
}

var (
	notes   []Note
	nextID  = 1
	tmpl    = template.Must(template.ParseFiles("templates/layout.html", "templates/index.html"))
	noteMux sync.Mutex
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	noteMux.Lock()
	defer noteMux.Unlock()
	tmpl.ExecuteTemplate(w, "layout.html", notes)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		content := r.FormValue("content")
		if title != "" && content != "" {
			noteMux.Lock()
			notes = append(notes, Note{ID: nextID, Title: title, Content: content})
			nextID++
			noteMux.Unlock()
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		idStr := r.FormValue("id")
		id, _ := strconv.Atoi(idStr)
		noteMux.Lock()
		for i := range notes {
			if notes[i].ID == id {
				notes = append(notes[:i], notes[i+1:]...)
				break
			}
		}
		noteMux.Unlock()
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
