package crawler

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/vivekmurali/spidey/pkg/db"
	"golang.org/x/net/html"
)

// A recursive function that runs till count counts down
func GetPage(ch chan *db.Data, url string, wg *sync.WaitGroup) {
	defer wg.Done()

	data, err := parsePage(url)
	if err != nil {
		log.Println(err)
		return
	}

	ch <- data
}

func parsePage(u string) (*db.Data, error) {
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("Bad status: %s", u)
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		return nil, err
	}

	// fmt.Println(getTitle(doc))

	b, err := body(doc)
	if err != nil {
		return nil, err
	}

	var links []string

	title := getTitle(doc)
	body := getBodyString(b)
	links = getLinks(b)

	var wg sync.WaitGroup
	wg.Add(1)

	go index(u, title+" "+body, &wg)

	for i, v := range links {
		if !isUrl(v) {
			links[i], err = absoluteURL(u, v)
			if err != nil {
				return nil, err
			}
		}
	}

	data := db.Data{URL: u, Title: title, Content: body, Links: links, Last_parsed: time.Now().Unix()}

	wg.Wait()
	// err = data.Insert()
	// if err != nil {
	// 	return err
	// }

	return &data, nil
}
