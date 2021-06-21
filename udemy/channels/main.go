package main

// concurrency is not parallelism

// with one cpu core we are using concurrency
// with multiple cpus we are running multiple threads at the same time

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com", 
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
		"http://stackoverflow.com",
	}

	for _, link := range links {
		go checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}
