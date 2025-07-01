## 📝 Go To-Do App

A simple, lightweight to-do list web application built using **Go**, **HTML templates**, and **basic CSS**. It allows users to add, mark as done, and delete tasks — all in-memory (no database).

---

### 🚀 Features

* Add new tasks
* Mark tasks as done
* Delete completed or pending tasks
* Fully responsive and styled UI
* Uses HTML templating (`html/template`) for server-side rendering

---

### 🛠 Tech Stack

* **Go** (net/http, html/template)
* **HTML/CSS**
* No external database – data is stored in memory

---

### 📁 Project Structure

```
go-todo/
├── main.go                  # Main application logic and routes
├── templates/
│   ├── layout.html          # Base HTML layout
│   └── index.html           # Main content template
├── static/
│   └── style.css            # Styling for the to-do list
```

---

### ⚙️ How to Run

1. **Clone the repo**

   ```bash
   git clone https://github.com/sa7vic/learn-go.git
   cd learn-go
   cd go-todo
   ```

2. **Run the app**

   ```bash
   go run main.go
   ```

3. **Open in your browser**

   ```
   http://localhost:8080
   ```

---

### 📚 Learning Highlights

* How to build a web server in Go
* Routing with `http.HandleFunc`
* Using `html/template` for rendering dynamic data
* Working with forms (POST/GET)
* Basic concurrency-safe data handling with `sync.Mutex`

---

### 📦 Future Improvements

* Add persistent storage (JSON, BoltDB, SQLite)
* User authentication
* Filtering (all, active, completed)
* RESTful API version

---

