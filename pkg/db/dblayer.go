package db

import (
	"errors"
	"strings"

	"github.com/mattn/go-sqlite3"
)

type Data struct {
	URL         string
	Title       string
	Content     string
	Links       []string
	Last_parsed int64
}

func (d Data) Insert() error {

	var e sqlite3.Error

	links := strings.Join(d.Links, ";")
	_, err := db.Exec("insert into documents (url, title, content, links, last_parsed) values($1, $2, $3, $4, $5)", d.URL, d.Title, d.Content, links, d.Last_parsed)
	if err != nil {
		if errors.As(err, &e) && e.ExtendedCode == sqlite3.ErrConstraintUnique {
			_, err := db.Exec("update documents set title = $1, content = $2, links = $3, last_parsed = $4 where url = $5", d.Title, d.Content, links, d.Last_parsed, d.URL)
			if err != nil {
				return err
			}

		} else if errors.As(err, &e) && e.ExtendedCode != sqlite3.ErrConstraintUnique {
			return err
		}
	}
	for _, v := range d.Links {
		_, err = db.Exec("insert into documents (url) values($1)", v)
		if err != nil {
			if errors.As(err, &e) {
				continue
			}
			return err
		}
	}
	return nil
}

func GetLinks() ([]string, error) {
	var links []string

	rows, err := db.Query("select url from documents where title is null")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var link string
		err = rows.Scan(&link)
		if err != nil {
			return nil, err
		}
		links = append(links, link)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return links, nil
}
