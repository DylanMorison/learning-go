package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)
}

	// fmt.Println(resp)
	// bs := make([]byte, 99999)
	// n, e := resp.Body.Read(bs)
	// fmt.Println(string(bs))
	// fmt.Println(n)
	// fmt.Println(e)