package utils

import (
	"fmt"

	"golang.org/x/net/html"
)

func ParsePaginationUrl(doc *html.Node) {
	var links []string

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	for _, link := range links {
		fmt.Println(link)
	}
}
