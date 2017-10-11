package inspection

import (
	"net/http"
	"errors"
	"golang.org/x/net/html"
	"io"
)

// scan website and returns unique links
type Scanner struct {
	links []string
}

// fetches website for given url.
func (scanner *Scanner) fetchWebsite(url string) (*io.ReadCloser, error) {
	response, _ := http.Get(url)

	if response.StatusCode == 200 {
		return &response.Body, nil
	}

	return nil, errors.New(response.Status)
}


func (scanner *Scanner) parseUrls(body *io.ReadCloser) []string {
	var links []string

	readBody := html.NewTokenizer(body)

	for {
		line := readBody.Next()
		switch {
			case line == html.ErrorToken:
				// End of the document, we're done
				return links
			case line == html.StartTagToken:
				text := readBody.Token()

				if text.Data == "a" {
					for _, attr := range text.Attr {
						if attr.Key == "href" {
							links = append(links, attr.Key)
						}
					}
				}
			}
		}
}

func (scanner *Scanner) ScanHost(url string) {

}
