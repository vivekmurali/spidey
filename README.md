# SPIDEY

A small search engine

# Flow

- Crawl command
- Open seed links and their following links
- For each page, remove stop words and stem them
- Inverse index them

# TODO

- [x]  Init should create DB file
- [ ]  Change viper location to new folder


# Schema

### links

ID, URL, title, content, links, last_parsed

### index

word, links

https://pkg.go.dev/github.com/cockroachdb/pebble