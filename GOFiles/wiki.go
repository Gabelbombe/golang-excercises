package main

import (
	"errors"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
	//	"fmt"
)

// data struct
type Page struct {
	Title string
	Body  []byte // were []byte is a "byte slice",
} // REF: http://blog.golang.org/go-slices-usage-and-internals

/**
 * Globals for caching
 */
var templates = template.Must(template.ParseFiles("template/edit.html", "template/view/html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/(\\w+)$") // ^/(edit|save|view)/([a-zA-Z0-9]+)$

/**
 * Save process
 */
func (p *Page) save() error {
	filename := "templates/" + p.Title + ".html"
	return ioutil.WriteFile(filename, p.Body, 0600) // 0600 permissions
}

/**
 * Page Loading
 */
func loadHandler(title string) (*Page, error) {
	filename := "template/" + title + ".html"
	body, _ := ioutil.ReadFile(filename)
	return &Page{Title: title, Body: body}, nil
}

/**
 * Get Title by substring matching
 */
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2].nil // were title is the second sub expression
}

/**
 * Template renderer: Has error handling
 */
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, "templates/"+tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/**
 * View Handler
 */
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadHandler(title)
	if err != nil {
		http.Redirect(w, r, "/edit"+title, http.StatusFound) // Adds 302 if not-found
		return
	}
	renderTemplate(w, "view", p)
}

/**
 * Edit Handler
 */
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadHandler(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

/**
 * Save Handler
 */
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

/**
 * Do work...
 */
func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	http.ListenAndServe(":8080", nil)
}
