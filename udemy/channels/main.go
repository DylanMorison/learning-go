package main

// concurrency is not parallelism

// with one cpu core we are using concurrency
// with multiple cpus we are running multiple threads at the same time

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
		"http://stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// this is a blocking line of code
	// fmt.Println(<- c)

	// for {
	// 	go checkLink(<- c, c)
	// }

	// watch channel c
	// whenever a value comes out to it, assign it to l
	// then body of for loop is instantly started
	// for l := range c {
	// 	time.Sleep(time.Second)  
	// 	go checkLink(l, c)
	// }

	for l := range c {
		go func () {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}()
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}