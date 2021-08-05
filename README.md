# SPIDEY

### A web crawler + search engine



1. Store a list of seed URLs in the database
2. Go through the seed URLs and store non duplicate hyperlinks in the database
3. Get all URLs from the database and continue crawling
4. Commands that I can use to search all the data I have crawled
5. Can reset stored data to run again with a different set of websites



# Procedure

* Store the webpage in the database
* Parse the HTML
* Get a list of hyperlinks
* Store the hyperlinks in the database if they don't already exist





# Database

* URL
* Time last parsed
* Meta information
* Content





# TODO

* [ ] Create a database
* [ ] Put the Seed URLs in the database ( write a method call to do this)
* [ ] Worker
  * [ ] Retrieve all URLs from the database
  * [ ] Spawn a new worker that parses the webpage
  * [ ] Store the hyperlinks in the database (unique links)
  * [ ] Store the content
* [ ] Time the process

