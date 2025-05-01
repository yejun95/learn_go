package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.soundcloud.com/",
		"https://www.facebook.com/",
	}
	for _, url := range urls {
		fmt.Println("Checking:", url)
		hitURL(url)
	}
}

func hitURL(url string) error {
	resp, err := http.Get(url)
	if err == nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}

	return nil
}
