package db

import (
	"fmt"
	"log"
	"strings"
)

type Data struct {
	URL         string
	Title       string
	Content     string
	Links       []string
	Last_parsed int64
}

func Insert(d []Data) {

	if len(d) < 1 {
		fmt.Println("DIdn't receive enough data")
		return
	}
	stmt := "insert or replace into documents (url, title, content, links, last_parsed) values %s"

	valueStrings := []string{}
	valueArgs := []interface{}{}
	linkStrings := []string{}
	linkArgs := []interface{}{}

	for _, v := range d {
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?)")
		links := strings.Join(v.Links, ";")

		valueArgs = append(valueArgs, v.URL)
		valueArgs = append(valueArgs, v.Title)
		valueArgs = append(valueArgs, v.Content)
		valueArgs = append(valueArgs, links)
		valueArgs = append(valueArgs, v.Last_parsed)

		for _, v1 := range v.Links {
			linkStrings = append(linkStrings, "(?)")
			linkArgs = append(linkArgs, v1)
		}
	}

	stmt = fmt.Sprintf(stmt, strings.Join(valueStrings, ","))
	// fmt.Println(stmt)

	_, err := db.Exec(stmt, valueArgs...)
	if err != nil {
		log.Fatalf("Couldn't insert into SQLite: %v", err)
	}

	stmt = "insert or ignore into documents (url) values %s"
	stmt = fmt.Sprintf(stmt, strings.Join(linkStrings, ","))

	_, err = db.Exec(stmt, linkArgs...)
	if err != nil {
		log.Fatalf("Couldn't insert into SQLite: %v", err)
	}
}

func GetLinks() ([]string, error) {
	var links []string

	rows, err := db.Query("select url from documents where title is null limit 50")
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
