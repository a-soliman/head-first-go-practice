package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type page struct {
	url  string
	size int
}

func resBodySize(url string, channel chan page) {
	fmt.Println("Getting ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- page{url: url, size: len(body)}
}

func main() {
	pages := make(chan page)
	urls := []string{
		"http://www.a-soliman.com",
		"https://example.com",
		"https://golang.org",
		"https://golang.org/doc",
		"https://nodejs.org",
	}

	for _, url := range urls {
		go resBodySize(url, pages)
	}

	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("%s: %d\n", page.url, page.size)
	}
}
