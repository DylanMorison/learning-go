package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Cubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "five", "six", "seven", "eight", "nine", "ten", "jack", "queen", "king"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// any var of type deck
// has access to the print method
// this is the point of a receiver method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) {
	fmt.Println(d)
}