package main

import "fmt"

// if there is any other type within this 
// program that has a function called getGreeting()
// that returns a string, that type is a member of type bot
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// getGreeting is associated with englishBot
func (englishBot) getGreeting() string {
	// very custom logic for generating an English greeting
	return "Hi There!"
}
 
// getGreeting is associated with spanishBot
func (spanishBot) getGreeting() string {
	// very custom logic for generating Spanish greeting
	return "Hola!"
}
