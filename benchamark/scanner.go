package benchamark

import (
	"net/http"
	"errors"
	"golang.org/x/net/html"
	"io"
	"strings"
)

// parser for website
type Parser struct{}

// parse website to find all urls.
func (parser *Parser) findUrls(body io.ReadCloser, url string) []string {
	var links []string

	readBody := html.NewTokenizer(body)

	for {
		line := readBody.Next()
		if line == html.StartTagToken {
			value := parser.getHrefValue(readBody.Token())

			// gets only internal urls.
			if (value != "") && strings.HasPrefix(value, "/")  {
				links = append(links, url + value)
			}
		}
		if line == html.ErrorToken {
			return links
		}
	}
}

// method returns value of href if exists.
func (parser *Parser) getHrefValue(text html.Token) string {
	if text.Data == "a" {
		for _, attr := range text.Attr {
			if attr.Key == "href" {
				return attr.Val
			}
		}
	}
	return ""
}

// scans website and returns unique links
type Scanner struct {
	links  []string
	parser Parser
}

// fetches website for given url.
func (scanner *Scanner) getBody(url string) (io.ReadCloser, error) {
	response, _ := http.Get(url)

	if response.StatusCode == 200 {
		return response.Body, nil
	}

	return nil, errors.New(response.Status)
}

// scans host to get all internal urls.
func (scanner *Scanner) ScanHost(url string) []string {
	body, _ := scanner.getBody(url)
    return scanner.parser.findUrls(body, url)
}
