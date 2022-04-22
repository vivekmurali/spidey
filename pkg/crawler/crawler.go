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
func GetPage(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	err := parsePage(url)
	if err != nil {
		log.Println(err)
	}
}

func parsePage(u string) error {
	res, err := http.Get(u)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return fmt.Errorf("Bad status: %s", u)
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		return err
	}

	// fmt.Println(getTitle(doc))

	b, err := body(doc)
	if err != nil {
		return err
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
				return err
			}
		}
	}

	data := db.Data{URL: u, Title: title, Content: body, Links: links, Last_parsed: time.Now().Unix()}

	wg.Wait()
	err = data.Insert()
	if err != nil {
		return err
	}

	return nil
}
