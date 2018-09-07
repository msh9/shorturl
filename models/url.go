package models

import (
	"net/url"
)

// URL defines a URL as received from a client requesting a short version
// of the given URL
type URL struct {
	url  *url.URL
	path string
}

// ToShort renders a non-unique, but determinitically repeatable short version of the URL
func (url *URL) ToShort() string {
	return ""
}
