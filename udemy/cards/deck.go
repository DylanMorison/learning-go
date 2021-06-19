package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
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

// this function has the reciever: (d deck)
// which allows us to call deck.toString()
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// This function reads byte slice from filename
// byte slice => string => []string => deck
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error a and return a call to newDeck()
		// Option #2 - Log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		// take whatever is at newPosition, and assign it to i
		// then take whatever is at i, and assign it to newPosition
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
