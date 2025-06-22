# Static Website with Form Handling in Go

This project is a simple static website built using HTML and CSS, with form submission handled by a Go backend.

## Features

- Static home page with links
- Form page to submit user information (name and address)
- Basic server-side form handling with Go
- A `/hello` endpoint that returns a greeting

## 📁 Project Structure

```text
/project-root
│
├── static/
│   ├── index.html         # Home page
│   ├── form.html          # Form page
│   └── styles.css         # CSS styles
│
├── main.go                # Go server handling routes and form submission
└── README.md              # Project documentation

````

## How It Works

- The Go server serves all files in the `./static` directory.
- Submitting the form sends a POST request to `/form`, which the server handles and responds with the submitted data.
- Visiting `/hello` returns a plain "Hello!" message.

## Running the Server

1. **Install Go** (if not already installed):  
   [https://golang.org/dl/](https://golang.org/dl/)

2. **Run the server**:

   ```bash
   go run main.go
   ```

3. **Open in browser**:
   Navigate to `http://localhost:8080` to view the website.

## Endpoints

* `/` - Static home page
* `/form.html` - Form page
* `/form` - POST handler for submitted form data
* `/hello` - Simple GET route returning a greeting

## Example Form Submission

Submitting this:

```
Name: John Doe
Address: 123 Main St
```

Will render:

```html
<h2>POST request successful</h2>
<p>Name: John Doe</p>
<p>Address: 123 Main St</p>
```

