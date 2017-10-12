package settings

import "strings"

// simple url normalize method. The last char of url must be "/".
func NormalizeUrl(url string) string {
	if strings.HasSuffix(url, "/") {
		return url
	}
	return url + "/"
}
