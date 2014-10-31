package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
//	"fmt"
)

// data struct
type Page struct {
	Title string
	Body  []byte									// were []byte is a "byte slice",
}													// REF: http://blog.golang.org/go-slices-usage-and-internals


/**
 * Save process
 */
func (p *Page) save() error {
	filename := "templates/" +  p.Title + ".html"
	return ioutil.WriteFile(filename, p.Body, 0600) // 0600 permissions
}


/**
 * Page Loading
 */
func loadPage (title string) (*Page, error) {
	filename  := title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	return &Page {Title: title, Body: body}, nil
}


/**
 * Template renderer: Has error handling
 */
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles("templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


/**
 * View Handler
 */
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title  := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit" + title, http.StatusFound) // Adds 302 if not-found
		return
	}
	renderTemplate(w, "view", p)
}


/**
 * Edit Handler
 */
func editHandler(w http.ResponseWriter, r *http.Request) {
	title  := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
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
	body  := r.FormValue("body")
	p     := &Page{Title: title, Body: []byte(body)}
	p.save()

	http.Redirect(w, r, "/view/" + title, http.StatusFound)
}


// run
func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	http.ListenAndServe(":8080", nil)
}
