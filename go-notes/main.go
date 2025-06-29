package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/gomarkdown/markdown"
)

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var (
	tmpl     = template.Must(template.New("").Funcs(template.FuncMap{
		"safeHTML": func(s string) template.HTML { return template.HTML(s) },
	}).ParseGlob("templates/*.html"))
	notes    = []Note{}
	noteMux  sync.Mutex
	nextID   = 1
	dataFile = "data/notes.json"
)

func main() {
	loadNotes()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/note/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/delete/", deleteHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", nil)
}


func loadNotes() {
	noteMux.Lock()
	defer noteMux.Unlock()

	data, err := ioutil.ReadFile(dataFile)
	if err == nil {
		json.Unmarshal(data, &notes)
		for _, note := range notes {
			if note.ID >= nextID {
				nextID = note.ID + 1
			}
		}
	} else {
		notes = []Note{}
	}
}

func saveNotes() {
	data, _ := json.MarshalIndent(notes, "", "  ")
	os.MkdirAll(filepath.Dir(dataFile), os.ModePerm)
	_ = ioutil.WriteFile(dataFile, data, 0644)
}


func indexHandler(w http.ResponseWriter, r *http.Request) {
	noteMux.Lock()
	defer noteMux.Unlock()
	tmpl.ExecuteTemplate(w, "layout.html", map[string]interface{}{
		"Content": "index",
		"Notes":   notes,
	})
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "layout.html", map[string]interface{}{
		"Content": "new",
	})
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		title := r.FormValue("title")
		content := r.FormValue("content")

		noteMux.Lock()
		notes = append(notes, Note{
			ID:      nextID,
			Title:   title,
			Content: content,
		})
		nextID++
		saveNotes()
		noteMux.Unlock()
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/note/")
	id, _ := strconv.Atoi(idStr)

	noteMux.Lock()
	defer noteMux.Unlock()
	for _, note := range notes {
		if note.ID == id {
			rendered := markdown.ToHTML([]byte(note.Content), nil, nil)
			tmpl.ExecuteTemplate(w, "layout.html", map[string]interface{}{
				"Content": "view",
				"Note":    note,
				"HTML":    string(rendered),
			})
			return
		}
	}
	http.NotFound(w, r)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/edit/")
	id, _ := strconv.Atoi(idStr)

	noteMux.Lock()
	defer noteMux.Unlock()
	for _, note := range notes {
		if note.ID == id {
			tmpl.ExecuteTemplate(w, "layout.html", map[string]interface{}{
				"Content": "edit",
				"Note":    note,
			})
			return
		}
	}
	http.NotFound(w, r)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		id, _ := strconv.Atoi(r.FormValue("id"))
		title := r.FormValue("title")
		content := r.FormValue("content")

		noteMux.Lock()
		for i := range notes {
			if notes[i].ID == id {
				notes[i].Title = title
				notes[i].Content = content
				break
			}
		}
		saveNotes()
		noteMux.Unlock()
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/delete/")
	id, _ := strconv.Atoi(idStr)

	noteMux.Lock()
	for i := range notes {
		if notes[i].ID == id {
			notes = append(notes[:i], notes[i+1:]...)
			break
		}
	}
	saveNotes()
	noteMux.Unlock()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
