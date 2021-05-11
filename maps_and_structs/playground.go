package main

import "fmt"

func main() {
	// map one key type to one value type
	// keys have to be able to be tested for equality
	// slices maps and functions cannot be tested for equality
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California": 214214525,
		"Texas":      2343243243,
		"Florida":    432432432,
	}
	fmt.Println(statePopulations)
}
