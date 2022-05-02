package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	bolt "go.etcd.io/bbolt"
)

var db *sql.DB
var KV *bolt.DB

func Initalize() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	path := filepath.Join(home, "spidey", "spidey.db")
	path += "?cache=shared&mode=rwc&_busy_timeout=9999999"
	db, err = sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}

	path = filepath.Join(home, "spidey", "kv.db")
	KV, err = bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

//TODO: TEST UNIQUE CREATE TABLE
func InitDB() {
	sqlStmt := `create table documents (id integer primary key, url string unique, title string, content string, links string, last_parsed integer);`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}

	stmt := `create table indexes (id integer primary key, word string, documents string);`
	_, err = db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}

	err = KV.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("bucket"))
		if err != nil {
			return err
		}
		return nil
	})
}
