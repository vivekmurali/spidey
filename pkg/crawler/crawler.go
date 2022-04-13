package crawler

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

// A recursive function that runs till count counts down
func GetPage(url string, wg *sync.WaitGroup, count int) {
	if count < 1 {
		wg.Done()
		return
	}
	links, err := parsePage(url)
	if err != nil {
		wg.Done()
		log.Println(err)
		return
	}

	var wg1 sync.WaitGroup
	for _, v := range links {
		// Convert relative URLs to absolute
		wg1.Add(1)
		GetPage(v, &wg1, count-1)
	}
	wg1.Wait()

	wg.Done()
}

func parsePage(url string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("Bad status")
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

	_ = b

	// store in db
	title := getTitle(doc)
	body := getBodyString(b)
	links := getLinks(b)
	_, _ = title, body
	fmt.Println(title)
	fmt.Println("")

	return links, nil
}

func body(doc *html.Node) (*html.Node, error) {
	var body *html.Node
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "body" {
			body = node
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	if body != nil {
		return body, nil
	}
	return nil, errors.New("Missing <body> in the node tree")
}

func getTitle(doc *html.Node) string {
	var title string
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "title" {
			if node.FirstChild.Data != "" {
				title = node.FirstChild.Data
			}
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	return title
}

func getBodyString(b *html.Node) string {
	var buf bytes.Buffer

	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.TextNode {
			buf.Write([]byte(node.Data))
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			if child.Type == html.ElementNode && child.Data == "script" {
				continue
			}
			crawler(child)
		}
	}
	crawler(b)
	return standardizeSpaces(buf.String())
	// return strings.TrimSpace(buf.String())
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// Check links
func getLinks(b *html.Node) []string {
	links := make([]string, 0, 0)

	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, a := range node.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
					break
				}
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(b)

	return links
}
