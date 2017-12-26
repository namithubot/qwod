package data

/*
Websites that don't require keys:
	-> https://urban-word-of-the-day.herokuapp.com/
	-> Option 2: Build your own (at least a proxy server)...can even collect from multiple sources and randomize
	-> Ask users to get their own keys to use
*/

import (
	"net/http"
)

func GetQuote() (*http.Response, error) {
	return http.Get("https://urban-word-of-the-day.herokuapp.com/")
}

func GetWord() (*http.Response, error) {
	return http.Get("https://urban-word-of-the-day.herokuapp.com/")
}
