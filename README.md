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

ID, URL, content, links, last_parsed

### index

ID, word, links