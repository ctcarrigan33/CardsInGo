package main

import "fmt"

// Create new type of deck
// which is a slice of strings

type deck []string

// receiver loops through deck of cards, prints values
// now, any var of type 'deck' has access to 'print'
// 'd' is reference to actual instance of the deck we're using
// here, 'd' points to result of cards.print() in main.go
// similar to 'this' in JavaScript
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}

}
