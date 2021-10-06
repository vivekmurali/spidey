# SPIDEY

### A web crawler + search engine

This tool can be used to search through pages that don't have a search option.


### Commands

* version
* download
* search
* clear

#### spidey

Initialize the application

`spidey`

#### version

Prints the version of the project

#### download

The crawler collects and stores pages

`spidey download -l URL`

##### options:

* -f input.yml {Use this file to set options and URLS}
* -fl input.txt {The file will contain a list of links that the crawler will download from}
* -s "URL" {Crawler will use the URL as the seed link}



#### search

Search for a term in the pages stored by the crawler

`spidey search search term`

#### clear

Clear all the crawled pages
