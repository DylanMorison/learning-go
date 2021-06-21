package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 10000)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}

// We want to read the contents of a file and print it out to the terminal
// (similar to http program)
// The file to open will be provided as a command line argument in the terminal.
// Example: `go run main.go myfile.txt` should print the contents of myfile.txt
// To read the arguments provided to a program, reference the variable `os.Args`,
// which is a silce of type `string`
// To open a file, check out docs for the `open` function in the `os` package
// What interface does the `File` type implement?
// If the `File` type implements the `Reader` interface, maybe use io.Copy function
