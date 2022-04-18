package crawler

import (
	"bytes"
	"errors"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

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

func absoluteURL(u, v string) (string, error) {
	base, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	rel, err := url.Parse(v)
	if err != nil {
		return "", err
	}

	return base.ResolveReference(rel).String(), nil
}

func isUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
