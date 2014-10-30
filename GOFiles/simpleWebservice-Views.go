package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

func genericHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I love %s...\n", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _  := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/", 	  genericHandler)
	http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
