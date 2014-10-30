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
	body, _ := ioutil.ReadFile(filename)			// _ is a blank identifier, essentially assigning error to /dev/null
	return &Page {Title: title, Body: body}, nil
}

func oopsHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I love %s...\n", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _  := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title  := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("/pages/edit.html")
	t.Execute(w, p)
}


func main() {
	http.HandleFunc("/", 	  oopsHandler)

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	http.ListenAndServe(":8080", nil)
}
