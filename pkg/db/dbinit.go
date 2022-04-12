package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	path := filepath.Join(home, "spidey", "spidey.db")

	db, err = sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}
}

func InitDB() {
	sqlStmt := `create table documents (id integer primary key, url string, title string, content string, links string, last_parsed integer);`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}

	stmt := `create table indexes (id integer primary key, word string, documents string);`
	_, err = db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}

}
