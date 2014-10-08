package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:8000", nil)
}

/**
 * w is an instantiation of http.ResponseWriter
 * r is an instantiation of http.Response (read)
 */
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, GOLang")
}

//RUNBOX: curl -X GET localhost:8000 -> Hello, GOLang
