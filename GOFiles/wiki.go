package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"fmt"
)

// data struct
type Page struct {
	Title string
	Body  []byte									// were []byte is a "byte slice",
}													// REF: http://blog.golang.org/go-slices-usage-and-internals

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600) // 0600 permissions
}

func loadPage (title string) (*Page, error) {
	filename  := title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	return &Page {Title: title, Body: body}, nil
}

func oopsHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I love men...")
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles("templates/" + tmpl + ".html")
	t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title  := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit" + title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title  := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}


func main() {
	http.HandleFunc("/", 	  oopsHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
//	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}
