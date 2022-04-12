# SPIDEY

A small search engine

# Flow

- Crawl command
- Open seed links and their following links
- For each page, remove stop words and stem them
- Inverse index them

# TODO

- [x] Set up sqlite
- [ ] Get links from page
- [ ] remove stop words
- [ ] Index
- [ ] Search
- [ ]  Change viper location to new folder


# Schema

### documents

ID, URL, title, content, links, last_parsed

### index

ID, word, links

https://pkg.go.dev/github.com/cockroachdb/pebble