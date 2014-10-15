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

	err := db.Exec("CREATE DATABASE api")
	if err != nil {
		log.Println(err)
	}

	err := dbExec("CREATE TABLE api.bookmarks")
	if err != nil {
		log.Println(err)
	}
}

func dropDb() {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_CONNECTION"))
	if err != nil {
		log.Fatal(err)
		return
	}

	err := db.Exec("DROP DATABASE api")
	if err != nil {
		log.Println(err)
	}


}
