package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

func main() {
	fmt.Printf("%v", a)
	fmt.Printf("%v", b)
	fmt.Printf("%v", c)
}
