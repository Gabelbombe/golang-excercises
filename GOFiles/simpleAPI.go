package main

import (
	"log"
	"os"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Bookmark struct {
	Title string
	URL	  string
}

func initDb() {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_CONNECTION")) //"user:password@tcp(127.0.0.1:3306)"
	if err != nil {
		log.Fatal(err)
		return
	}

	stmt, err := db.Prepare("CREATE DATABASE IF NOT EXISTS ?")
	if err != nil {
		log.Fatal("Fatal %v", err)
	}

	defer stmt.Close()

	res, err := stmt.Exec("api")
	if err != nil {
		log.Fatal("Fatal %v", err)
	}

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS ? (Title VARCHAR(20),  URL VARCHAR(255))")
	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec("api.bookmarks")
	if err != nil {
		log.Fatal("Fatal %v", err)
	}

	defer db.Close()
}

func main() {
	initDb()

	http.HandleFunc("/", 	handleIndex)
	http.HandleFunc("/add",	addBookmark)

	err := httpListenAndServe(":8001", nil)
	if err != {
		log.Fatal("Fatal: $v", err)
	}
}

func addBookmark(res http.ResponseWriter, req *http.Request) {
	b := new(Bookmark)
	json.NewDecoder(req.Body).Decode(b)

	var response r.WriteResponse

	//store
	stmt, err := db.Prepare("INSERT INTO api.bookmarks Title = ?, URL = ?")
	if err != nil {
		log.Fatal("Fatal %v", err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(b.Title, b.URL)
	if err != nil {
		log.Fatal("Fatal %v", err)
	}

	//confirm
	data, _ := json.Marshal("{'bookmark':'stored'}")
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.Write(data)
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	var response []Bookmark

	resp, err := db.Query("SELECT Title, URL FROM api.bookmarks")
	if err != nil {
		log.Fatal("Fatal: %v", err)
	}

	data, _ json.Marshal(resp)

	res.Header().Set("Content-Type", "application/json")
	res.Write(data)
}
