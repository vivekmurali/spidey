package crawler

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

func GetPage(url string, wg *sync.WaitGroup, count int) {
	// fmt.Println(url, count)
	parsePage(url)

	// Create new waitgroup before sending all links here
	// Recursive function

	wg.Done()
}

func parsePage(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		panic(err)
	}

	// fmt.Println(getTitle(doc))

	b, err := body(doc)
	if err != nil {
		panic(err)
	}

	_ = b
	// fmt.Println(getBodyString(b))
	fmt.Println("")

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
			title = node.FirstChild.Data
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
