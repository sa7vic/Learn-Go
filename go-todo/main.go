package main

import (
	"html/template"
	"net/http"
	"strconv"
	"sync"
)

type Todo struct {
	ID    int
	Title string
	Done  bool
}

var (
	todos   = []Todo{}
	nextID  = 1
	tmpl    = template.Must(template.ParseFiles("templates/layout.html", "templates/index.html"))
	todoMux sync.Mutex
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/done", doneHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer((http.Dir("static")))))

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	todoMux.Lock()
	defer todoMux.Unlock()

	tmpl.ExecuteTemplate(w, "layout.html", todos)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		title := r.FormValue("title")
		if title != "" {
			todoMux.Lock()
			todos = append(todos, Todo{ID: nextID, Title: title})
			nextID++
			todoMux.Unlock()
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func doneHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		idStr := r.FormValue("id")
		id, _ := strconv.Atoi(idStr)
		todoMux.Lock()
		for i := range todos {
			if todos[i].ID == id {
				todos[i].Done = true
				break
			}
		}
		todoMux.Unlock()
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		idStr := r.FormValue("id")
		id, _ := strconv.Atoi(idStr)
		todoMux.Lock()
		for i := range todos {
			if todos[i].ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				break
			}
		}
		todoMux.Unlock()
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
