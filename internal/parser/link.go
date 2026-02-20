package parser

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func LinkParser(n *html.Node, url string) {
	if n.Type == html.ElementNode {
		// baseUrl, err := GetBaseUrl(url)
		// if err != nil {
		// 	log.Fatal("Failed to parse url :", err)
		// }

		for _, attr := range n.Attr {
			if attr.Key == "href" {
				fmt.Println(attr.Val)
				// rawUrl := strings.Split(attr.Val, "/")
				// if j == 99 {
				// 	break
				// }
				// if rawUrl[0] == "http:" || rawUrl[0] == "https:" {
				// 	fmt.Println(attr.Val)
				// } else if attr.Val[0] == '/' {
				// 	fmt.Println(baseUrl + string(attr.Val))
				// }
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		LinkParser(c, url)
	}
}

func GetBaseUrl(givenUrl string) (string, error) {
	sptUrl := strings.Split(givenUrl, "/")
	baseUrl := ""
	if sptUrl[0] == "http:" || sptUrl[0] == "https:" {
		baseUrl = sptUrl[0] + "//" + sptUrl[2]
	} else {
		return "", errors.New("Invalid base url!")
	}
	return baseUrl, nil
}
