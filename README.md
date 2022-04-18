# SPIDEY

[![Go 1.18](https://github.com/vivekmurali/spidey/actions/workflows/ci.yml/badge.svg)](https://github.com/vivekmurali/spidey/actions/workflows/ci.yml)

A small search engine

# Flow

- Crawl command
- Open seed links and their following links
- For each page, remove stop words and stem them
- Inverse index them

# TODO

- [x] Set up sqlite
- [x] Get links from page
- [x] Set up insert into database and get from database
- [x] Change crawl.go to get links from db and crawl them instead of seed.txt
- [x] Change from recursion to checking database
- [x] Change github action workflow name
- [x] Make relative links absolute
- [x] remove stop words
- [ ] Index
- [ ] Search
- [ ] Change viper location to new folder
- [x] Database is locked


# Schema

### documents

ID, URL, title, content, links, last_parsed

### index

ID, word, links

https://pkg.go.dev/github.com/cockroachdb/pebble