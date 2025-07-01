# 📝 Go Notes

A simple web-based note-taking app built with **Go**. Users can add and delete notes, all managed in memory using Go's native data structures and concurrency safety via mutex locking.

---

## 🚀 Features

* Add and delete notes through a clean, responsive web interface.
* Simple in-memory storage (no database required).
* Styled with CSS for a polished UI.
* Concurrency-safe using Go's `sync.Mutex`.

---

## 🗂️ Project Structure

```
go-notes/
├── main.go                 # Main Go application
├── README.md               # Project readme
├── static/
│   └── style.css           # CSS styles
└── templates/
    ├── index.html          # Main content template
    └── layout.html         # Base HTML layout
```

---

## 🛠️ Requirements

* [Go](https://golang.org/dl/) 1.16 or later

---

## 🧪 Running the App

1. **Clone the repo**:

   ```bash
   git clone https://github.com/sa7vic/go-notes.git
   cd go-notes
   ```

2. **Run the server**:

   ```bash
   go run main.go
   ```

3. **Visit in your browser**:

   ```
   http://localhost:8080
   ```

---

## ✏️ Usage

* **Add a note**: Fill in the title and content fields and click "Add Note".
* **Delete a note**: Click the red "Delete" button next to any existing note.

---

## ⚠️ Notes

* This app does **not** persist data — all notes are stored in memory and will be lost when the server stops.
* To persist notes, consider integrating a database like SQLite or Postgres.

---

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

---

## 🙋‍♀️ Author

Made with 💙 in Go.

