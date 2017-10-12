package benchamark

import "strings"

// simple url normalize method. The last char of url must be "/".
func normalizeUrl(url string) string {
	if strings.HasSuffix(url, "/") {
		return url
	}
	return url + "/"
}
